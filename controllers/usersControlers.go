package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vineet12344/jwtgolang/initializers"
	"github.com/vineet12344/jwtgolang/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUP(c *gin.Context) {
	var Body struct {
		Email    string
		Password string
	}

	err := c.Bind(&Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Failed to read body",
		})

		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(Body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "password hashing errror",
		})
	}

	user := models.User{Email: Body.Email, Password: string(hashed)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ERROR": "Error while inserting inro DB",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Success": "User Created Sucessfully",
	})

}

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	err := c.Bind(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invald json input",
		})

		return
	}

	var user models.User

	initializers.DB.First(&user, "email=?", body.Email)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User email not found! Try SignIn first",
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Password",
		})

		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 10).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":       "Failed to create the Token",
			"description": err.Error(),
		})

		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	// c.SetCookie("Authorization", tokenString, 3600*24*10, "", "", false, false)
	c.SetCookie("Authorization", tokenString, 3600*24*10, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"success": "Token Generated",
	})

}

func Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Login-Success": "I am logged IN",
	})
}

package main

import (
	"fmt"

	"github.com/vineet12344/jwtgolang/controllers"
	"github.com/vineet12344/jwtgolang/initializers"
	"github.com/vineet12344/jwtgolang/middleware"
	"github.com/vineet12344/jwtgolang/router"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
	initializers.SyncDB()
}

func main() {

	fmt.Println("HELLO WORLD")

	r := router.InitRouter()

	r.POST("/signup", controllers.SignUP)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequiredAuth, controllers.Validate)

	r.Run(":8080")

}

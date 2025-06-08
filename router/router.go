package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(gc *gin.Context) {
		gc.JSON(http.StatusOK, gin.H{
			"message": "<h1>HELLO WORLD</h1>",
		})
	})

	return router
}

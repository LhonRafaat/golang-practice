package router

import (
	"github.com/gin-gonic/gin"
)


func SetupRouter() *gin.Engine{


	router:=gin.Default()

	router.GET("/notes", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	return router
}
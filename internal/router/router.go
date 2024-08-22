package router

import (
	"github.com/gin-gonic/gin"
)


func SetupRouter() *gin.Engine{


	router:=gin.Default()

	router.GET("/notes", func(c *gin.Context) {})

	return router
}
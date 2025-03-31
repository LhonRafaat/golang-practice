package routes

import "github.com/gin-gonic/gin"


func RegisterNoteRoutes(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"result": []string{"nots"},
			"isSuccess": true,
		})
	})
}
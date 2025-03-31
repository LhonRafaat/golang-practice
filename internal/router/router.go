package router

import (
	"lhon/postgres-rest/internal/router/routes"

	"github.com/gin-gonic/gin"
)


func SetupRouter() *gin.Engine{


	router:=gin.Default()

	noteRoutes:=router.Group("/notes")

	routes.RegisterNoteRoutes(noteRoutes)

	return router
}
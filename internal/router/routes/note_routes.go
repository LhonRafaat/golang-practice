package routes

import (
	"lhon/postgres-rest/internal/handler"
	"lhon/postgres-rest/internal/repository"
	"lhon/postgres-rest/internal/services"

	"github.com/gin-gonic/gin"
)


func RegisterNoteRoutes(router *gin.RouterGroup) {
	  noteRepo := repository.NewNoteRepository(repository.DB)
    noteService := services.NewNoteService(noteRepo)
    noteHandler := handler.NewNoteHandler(noteService)

	router.GET("/",noteHandler.GetAllNotes)
}
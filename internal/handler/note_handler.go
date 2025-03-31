package handler

import (
	"lhon/postgres-rest/internal/services"

	"github.com/gin-gonic/gin"
)


type NoteHandler struct{ 

	Service *services.NoteService
}

func NewNoteHandler(service *services.NoteService) *NoteHandler {
	return &NoteHandler{Service: service}
}


func (h *NoteHandler) GetAllNotes(c *gin.Context) {
	notes, err := h.Service.GetAllNotes()

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"result": notes,
		"isSuccess": true,
	})

}
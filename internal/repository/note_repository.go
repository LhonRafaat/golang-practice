package repository

import (
	"lhon/postgres-rest/internal/models"

	"gorm.io/gorm"
)



type NoteRepository struct{ 
	DB *gorm.DB
	
}


func NewNoteRepository(db *gorm.DB) *NoteRepository {
	return &NoteRepository{DB: db}
}

func (r *NoteRepository) GetAllNotes() ([]models.Note, error) {
	var notes []models.Note
	err:= r.DB.Find(&notes).Error
	return notes, err
}
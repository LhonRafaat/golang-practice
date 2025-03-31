package services

import (
	"lhon/postgres-rest/internal/models"
	"lhon/postgres-rest/internal/repository"
)


type NoteService struct { 

	Repo *repository.NoteRepository
}


func NewNoteService(repo *repository.NoteRepository) *NoteService {
	return &NoteService{Repo: repo}
}

func (s *NoteService) GetAllNotes() ([]models.Note, error) {
    return s.Repo.GetAllNotes()
}
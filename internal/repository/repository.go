package repository

import (
	"first/internal/domain"
	"github.com/jmoiron/sqlx"
)

type Lessons interface {
	GetLessonByID(lessonID int) (domain.Lesson, error)
}

type Repository struct {
	Lessons
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Lessons: NewLessonPostgres(db),
	}
}

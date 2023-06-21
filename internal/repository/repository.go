package repository

import (
	"first/internal/domain"
	"first/internal/repository/postgres"
	"first/internal/repository/storage"
	"github.com/jmoiron/sqlx"
)

type Lessons interface {
	GetLessonByID(lessonID int) (domain.Lesson, error)
}

type Videos interface {
	GetVideoByLessonID(videoID int) (domain.Video, error)
}

type Repository struct {
	Lessons
	Videos
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Lessons: postgres.NewLessonPostgres(db),
		Videos:  storage.NewVideoStorage("C:\\Users\\shokhrukh.davlatmama\\vidoes", postgres.NewVideoPostgres(db)),
	}
}

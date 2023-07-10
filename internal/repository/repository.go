package repository

import (
	"first/internal/domain"
	"first/internal/repository/postgres"
	"first/internal/repository/storage"
	"github.com/jmoiron/sqlx"
)

type Permission interface {
	GetUserPermissionForEndpoint(userID int) ([]string, error)
}

type Lessons interface {
	GetLessonByID(lessonID int) (domain.Lesson, error)
	GetAccessibleLessonsForUser(userID int) ([]string, error)
}

type Videos interface {
	GetVideoByLessonID(videoID int) (domain.Video, error)
}

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GetUser(username, password string) (domain.User, error)
}

type Repository struct {
	Lessons
	Videos
	Authorization
	Permission
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Lessons:       postgres.NewLessonPostgres(db),
		Videos:        storage.NewVideoStorage("C:\\Users\\shokhrukh.davlatmama\\vidoes", postgres.NewVideoPostgres(db)),
		Authorization: postgres.NewAuthPostgres(db),
		Permission:    postgres.NewPermissionPostgres(db),
	}
}

package service

import (
	"first/internal/domain"
	"first/internal/repository"
)

type Permission interface {
	GetUserPermissionForEndpoint(userID int, urlPath string) (bool, error)
}

type Lesson interface {
	GetLessonByID(userID, lessonID int, urlPath string) (domain.Lesson, error)
	IsLessonAccessibleForUser(userID, lessonID int) (bool, error)
}

type Video interface {
	GetVideoByLessonID(lessonID int) (domain.Video, error)
}

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GenerateToken(username, password string) (string, int, error)
	ParseToken(accessToken string) (int, error)
}

type Service struct {
	Lesson
	Video
	Authorization
	Permission
}

func NewService(repo *repository.Repository) *Service {
	return &Service{Lesson: NewLessonService(repo.Lessons),
		Video:         NewVideoService(repo.Videos),
		Authorization: NewAuthService(repo.Authorization),
		Permission:    NewPermissionService(repo.Permission),
	}
}

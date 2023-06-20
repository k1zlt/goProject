package service

import (
	"first/internal/domain"
	"first/internal/repository"
)

type Lesson interface {
	GetLessonByID(lessonID int) (domain.Lesson, error)
	//GetVideoByID(videoID int)
}

type Service struct {
	Lesson
}

func NewService(repo *repository.Repository) *Service {
	return &Service{Lesson: NewLessonService(repo.Lessons)}
}

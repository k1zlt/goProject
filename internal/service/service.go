package service

import (
	"first/internal/domain"
	"first/internal/repository"
)

type Lesson interface {
	GetLessonByID(lessonID int) (domain.Lesson, error)
	//GetVideoByID(videoID int)
}

type Video interface {
	GetVideoByLessonID(lessonID int) (domain.Video, error)
}

type Service struct {
	Lesson
	Video
}

func NewService(repo *repository.Repository) *Service {
	return &Service{Lesson: NewLessonService(repo.Lessons), Video: NewVideoService(repo.Videos)}
}

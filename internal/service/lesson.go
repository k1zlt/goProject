package service

import (
	"first/internal/domain"
	"first/internal/repository"
)

type LessonService struct {
	repo repository.Lessons
}

func NewLessonService(repo repository.Lessons) *LessonService {
	return &LessonService{repo: repo}
}

func (s *LessonService) GetLessonByID(lessonID int) (domain.Lesson, error) {
	return s.repo.GetLessonByID(lessonID)
}

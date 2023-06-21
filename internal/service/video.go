package service

import (
	"first/internal/domain"
	"first/internal/repository"
)

type VideoService struct {
	repo repository.Videos
}

func NewVideoService(repo repository.Videos) *VideoService {
	return &VideoService{repo: repo}
}

func (s *VideoService) GetVideoByLessonID(lessonID int) (domain.Video, error) {
	return s.repo.GetVideoByLessonID(lessonID)
}

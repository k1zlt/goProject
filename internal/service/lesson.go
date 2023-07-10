package service

import (
	"first/internal/domain"
	"first/internal/repository"
	"fmt"
	"strconv"
)

type LessonService struct {
	repo repository.Lessons
}

func NewLessonService(repo repository.Lessons) *LessonService {
	return &LessonService{repo: repo}
}

func (s *LessonService) GetLessonByID(userID, lessonID int, urlPath string) (domain.Lesson, error) {
	return s.repo.GetLessonByID(lessonID)
}

// IsLessonAccessibleForUser checks if the lesson with the provided lesson ID is accessible for the given user ID.
// It retrieves the accessible lesson IDs for the user from the repository and compares them to the provided lesson ID.
//
// Parameters:
// - userID: The ID of the user.
// - lessonID: The ID of the lesson to check accessibility for.
//
// Returns:
// - A boolean value indicating whether the lesson is accessible for the user.
// - An error if there was an issue retrieving the accessible lesson IDs or if the ID format is invalid.
func (s *LessonService) IsLessonAccessibleForUser(userID, lessonID int) (bool, error) {
	lessonsId, err := s.repo.GetAccessibleLessonsForUser(userID)
	if err != nil {
		return false, err
	}

	for _, id := range lessonsId {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return false, fmt.Errorf("invalid format of id: %v", err)
		}
		if idInt == lessonID {
			return true, nil
		}
	}

	return false, nil
}

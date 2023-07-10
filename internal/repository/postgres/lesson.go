package postgres

import (
	"first/internal/domain"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type LessonPostgres struct {
	db *sqlx.DB
}

func NewLessonPostgres(db *sqlx.DB) *LessonPostgres {
	return &LessonPostgres{
		db: db,
	}
}

func (r *LessonPostgres) GetLessonByID(lessonID int) (domain.Lesson, error) {
	var lesson domain.Lesson
	query := `
		SELECT lesson_id, content, title, video
		FROM lesson.lesson
		WHERE lesson_id = $1
	`
	if err := r.db.Get(&lesson, query, lessonID); err != nil {
		return domain.Lesson{}, fmt.Errorf("lesson does not exist: %v", err)
	}
	return lesson, nil
}

// GetAccessibleLessonsForUser retrieves the lesson IDs that are accessible for the user with the provided userID.
// It queries the database to fetch the lesson IDs associated with the user's accessibility.
//
// Parameters:
// - userID: The ID of the user.
//
// Returns:
// - A slice of strings representing the accessible lesson IDs.
// - An error if there was an issue querying the database or retrieving the accessible lesson IDs.
func (r *LessonPostgres) GetAccessibleLessonsForUser(userID int) ([]string, error) {
	var lessonsId []string
	query := `SELECT lesson_id
			FROM lesson.student_lesson
			WHERE student_id = $1`
	if err := r.db.Select(&lessonsId, query, userID); err != nil {
		return []string{}, fmt.Errorf("error while checking lesson permission: %v", err)
	}
	return lessonsId, nil
}

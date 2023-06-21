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
	query := fmt.Sprintf("SELECT lesson_id, content, title, video FROM %s WHERE lesson_id = $1", "lesson.lesson")

	if err := r.db.Get(&lesson, query, lessonID); err != nil {
		return lesson, err
	}

	return lesson, nil
}

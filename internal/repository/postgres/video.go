package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type VideoPostgres struct {
	db *sqlx.DB
}

func NewVideoPostgres(db *sqlx.DB) *VideoPostgres {
	return &VideoPostgres{db: db}
}

func (r *VideoPostgres) GetVideoFilename(lessonID int) (string, error) {
	var filename string
	query := fmt.Sprintf("SELECT video_filename FROM %s WHERE lesson_id = $1", "lesson.video")
	if err := r.db.Get(&filename, query, lessonID); err != nil {
		return "", err
	}

	return filename, nil
}

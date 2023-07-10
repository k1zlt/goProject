package domain

type Lesson struct {
	ID          int    `json:"lesson_id" db:"lesson_id"`
	Title       string `json:"title"`
	ContentText string `json:"contentText" db:"content"`
	Video       string `json:"video"`
}

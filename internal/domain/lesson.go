package domain

type Lesson struct {
	ID      int    `json:"lesson_id" db:"lesson_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Video   string `json:"video"`
}

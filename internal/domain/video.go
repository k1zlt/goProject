package domain

type Video struct {
	LessonID int    `json:"lesson_id"`
	Data     []byte `json:"data"`
}

package storage

import (
	"first/internal/domain"
	"first/internal/repository/postgres"
	"io/ioutil"
	"os"
	"path/filepath"
)

type VideoStorage struct {
	basePath      string
	postgresVideo *postgres.VideoPostgres
}

func NewVideoStorage(basePath string, postgresVideo *postgres.VideoPostgres) *VideoStorage {
	return &VideoStorage{basePath: basePath, postgresVideo: postgresVideo}
}

func (s *VideoStorage) GetVideoByLessonID(lessonID int) (domain.Video, error) {
	filename, err := s.postgresVideo.GetVideoFilename(lessonID)
	if err != nil {
		return domain.Video{}, err
	}

	videoPath := filepath.Join(s.basePath, filename)
	videoData, err := readVideoFile(videoPath)
	if err != nil {
		return domain.Video{}, err
	}

	video := domain.Video{
		LessonID: lessonID,
		Data:     videoData,
	}

	return video, nil
}

func readVideoFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	videoData, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return videoData, nil
}

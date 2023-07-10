package service

import (
	"first/internal/repository"
	"regexp"
)

type PermissionService struct {
	repo repository.Permission
}

func NewPermissionService(repo repository.Permission) *PermissionService {
	return &PermissionService{repo: repo}
}

// GetUserPermissionForEndpoint checks if the user with the provided userID has permission to access the specified URL path.
// It retrieves the endpoints associated with the user's permissions from the repository and matches them against the provided URL path.
//
// Parameters:
// - userID: The ID of the user.
// - urlPath: The URL path to check permission for.
//
// Returns:
// - A boolean value indicating whether the user has permission to access the specified URL path.
// - An error if there was an issue retrieving the user's endpoints or an error occurred during the matching process.
func (s *PermissionService) GetUserPermissionForEndpoint(userID int, urlPath string) (bool, error) {
	endpoints, err := s.repo.GetUserPermissionForEndpoint(userID)
	if err != nil {
		return false, err
	}

	for _, endpoint := range endpoints {
		if regexp.MustCompile(endpoint).MatchString(urlPath) {
			return true, nil
		}
	}

	return false, nil
}

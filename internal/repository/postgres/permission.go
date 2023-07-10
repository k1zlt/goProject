package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type PermissionPostgres struct {
	db *sqlx.DB
}

func NewPermissionPostgres(db *sqlx.DB) *PermissionPostgres {
	return &PermissionPostgres{db: db}
}

// GetUserPermissionForEndpoint retrieves the endpoints associated with the user's permissions for the specified userID.
// It queries the database to fetch the URLs from the permission table based on the user's role.
//
// Parameters:
// - userID: The ID of the user.
//
// Returns:
// - A slice of strings representing the endpoints associated with the user's permissions.
// - An error if there was an issue querying the database or retrieving the endpoints.
func (r *PermissionPostgres) GetUserPermissionForEndpoint(userID int) ([]string, error) {
	var endpoints []string
	query := `SELECT p.url
			FROM lesson.user AS u
			RIGHT JOIN lesson.permission AS p ON u.role_id = p.role_id
			WHERE u.id = $1;
			`
	if err := r.db.Select(&endpoints, query, userID); err != nil {
		return endpoints, fmt.Errorf("error while retrieving endpoints: %v", err)
	}
	return endpoints, nil
}

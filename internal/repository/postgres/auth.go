package postgres

import (
	"first/internal/domain"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user domain.User) (int, error) {

	var userID int
	queryForRole := `SELECT role_id FROM lesson.role WHERE name = $1`
	var roleID int
	if err := r.db.Get(&roleID, queryForRole, user.Role); err != nil {
		return 0, fmt.Errorf("role is invalid: %v", err)
	}

	query := fmt.Sprintf("INSERT INTO %s (username, password_hash, role_id) values ($1, $2, $3) RETURNING id", "lesson.user")
	row := r.db.QueryRow(query, user.Username, user.Password, roleID)

	err := row.Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (r *AuthPostgres) GetUser(username, password string) (domain.User, error) {
	var user domain.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", "lesson.user")
	if err := r.db.Get(&user, query, username, password); err != nil {
		return domain.User{}, err
	}
	return user, nil
}

package service

import (
	"crypto/sha1"
	"errors"
	"first/internal/domain"
	"first/internal/repository"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	salt       = "f4fdklgl"
	signingKey = "fjwdiof43fdv"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user domain.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

// tokenClaims is a custom struct that extends the jwt.StandardClaims struct to include a UserId field.
// It represents the claims contained within a JWT used for authentication and authorization.
type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

// GenerateToken generates a new JWT token for the provided username and password combination.
// It retrieves the user from the repository based on the username and hashed password,
// and creates a token with a set expiration time and the user's ID as a claim.
//
// Parameters:
// - username: The username of the user.
// - password: The password of the user.
//
// Returns:
// - The generated JWT token as a string.
// - The user_id
// - An error if there was an issue generating the token or retrieving the user from the repository.
func (s *AuthService) GenerateToken(username, password string) (string, int, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", 0, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{ExpiresAt: time.Now().Add(5 * time.Minute).Unix()},
		user.ID,
	})

	signedToken, err := token.SignedString([]byte(signingKey))
	return signedToken, user.ID, err
}

// ParseToken parses the provided access token and extracts the user ID from the token claims.
// It verifies the token's signature, checks the signing method, and returns the user ID associated with the token.
//
// Parameters:
// - accessToken: The access token string to be parsed.
//
// Returns:
// - The user ID extracted from the token claims.
// - An error if there was an issue parsing the token or the token claims are not of type *tokenClaims.
func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid singing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

// generatePasswordHash hashes the provided password using SHA-1 algorithm with a salt value.
// This function is used to securely store passwords in the database by converting them into a hashed representation.
//
// Parameters:
// - password: The password string to be hashed.
//
// Returns:
// - The hashed representation of the password as a string.
func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

package models

import (
	"database/sql"
	"strings"
	"time"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

type UserModel struct {
	DB *sql.DB
}

// To add a new record to the users table.
func (m *UserModel) Insert(name, email, password string) error {
	if strings.TrimSpace(name) == "" {
		return ErrInvalidName
	}

	email = strings.TrimSpace(email)
	if email == "" || !strings.Contains(email, "@") {
		return ErrInvalidEmail
	}

	if len(password) < 8 {
		return ErrInvalidPassword
	}

	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 60)
	if err != nil {
		return err
	}

	sqlStatement := `INSERT INTO users (name, email, hashed_password, created)
			VALUES($1, $2, $3, NOW())`

	_, err = m.DB.Exec(sqlStatement, name, email, HashedPassword)
	if err != nil {
		if strings.Contains(err.Error(), "users_uc_email") {
			return ErrDuplicateEmail
		}
	}
	return err
}

// Authenticate to verify whether a user exists with the provided email address and password and return the relevant
// user ID if they do.
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// Check if a user exists with a specific ID.
func (m *UserModel) Exists(id int) (bool, error) {
	return false, nil
}

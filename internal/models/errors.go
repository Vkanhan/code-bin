package models

import (
	"errors"
)

var (
	ErrNoRecords       = errors.New("models: no matching record found")
	ErrDuplicateEmail  = errors.New("email already exists")
	ErrInvalidName     = errors.New("name cannot be empty")
	ErrInvalidEmail    = errors.New("invalid email format")
	ErrInvalidPassword = errors.New("password must be at least 8 characters")
)

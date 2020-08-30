package errors

import "errors"

var (
	// ErrUserProfileNotFound is thrown when a UserProfile does not exists
	ErrUserProfileNotFound = errors.New("UserProfile not found")

	// ErrUserProfileAlreadyExists is thrown when a UserProfile exists in the repository
	ErrUserProfileAlreadyExists = errors.New("UserProfile already exists")

	// ErrInternalServerError is thrown when an unknown error is got
	ErrInternalServerError = errors.New("Unexpected error")
)

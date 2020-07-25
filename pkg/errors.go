package pkg

import "errors"

//noinspection ALL
var (
	ErrNotFound     = errors.New("Error: Document not found")
	ErrNoContent    = errors.New("Error: Document not found")
	ErrInvalidSlug  = errors.New("Error: Invalid slug")
	ErrExists       = errors.New("Error: Document already exists")
	ErrDatabase     = errors.New("Error: Database error")
	ErrUnauthorized = errors.New("Error: You are not allowed to perform this action")
	ErrForbidden    = errors.New("Error: Access to this resource is forbidden")
	ErrEmail        = errors.New("Error: Email not valid")
	ErrPassword     = errors.New("Error: Password must be greater than 6 chars")
	ErrNotAllowed   = errors.New("Error: Not allowed")
)

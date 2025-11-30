package errorx

import "errors"

// env.go
var (
	ErrTgTokenEmpty     = errors.New("TG_TOKEN is empty")
	ErrPostgresDSNEmpty = errors.New("POSTGRES_DSN is empty")
)

// reg.go
var (
	ErrAlreadyRegistered = errors.New("user is already registered")
	ErrUserNotFound      = errors.New("user not found")
)

// photo.go
var (
	ErrPhotoNotFound  = errors.New("photo not found")
	ErrPhotoDuplicate = errors.New("photo is duplicate")
)

// event/photo.go
var (
	ErrPhotoRequired     = errors.New("photo is required")
	ErrUserRequired      = errors.New("user is required")
	ErrFileIDRequired    = errors.New("file ID is required")
	ErrUniqueIDRequired  = errors.New("unique ID is required")
	ErrFileURLRequired   = errors.New("file URL is required")
	ErrCreatedAtRequired = errors.New("created_at is required")
)

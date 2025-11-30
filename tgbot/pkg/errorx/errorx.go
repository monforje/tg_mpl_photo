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
	ErrPhotoDuplicate = errors.New("photo is duplicate")
)

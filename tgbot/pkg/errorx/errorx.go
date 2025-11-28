package errorx

import "errors"

// env.go
var (
	ErrTgTokenEmpty     = errors.New("TG_TOKEN is empty")
	ErrPostgresDSNEmpty = errors.New("POSTGRES_DSN is empty")
)

package usecase

import "errors"

var (
	ErrAlreadyWatching = errors.New("already watching")
	ErrStreamNotFound  = errors.New("stream not found")
)

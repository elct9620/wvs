package usecase

import "context"

type Stream interface {
	Publish(event any) error
}

type StreamRepository interface {
	Find(context.Context, string) (Stream, error)
}

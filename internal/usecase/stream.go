package usecase

type Stream interface {
	Publish(event any) error
}

type StreamRepository interface {
	Find(string) (Stream, error)
}

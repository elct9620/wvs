package usecase

type Stream interface {
	Publish(event any) error
}

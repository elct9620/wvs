package subscriber

type DatabaseChange[T any] struct {
	Before *T
	After  *T
}

func (c *DatabaseChange[T]) Created() bool {
	return c.Before == nil && c.After != nil
}

func (c *DatabaseChange[T]) Updated() bool {
	return c.Before != nil && c.After != nil
}

func (c *DatabaseChange[T]) Deleted() bool {
	return c.Before != nil && c.After == nil
}

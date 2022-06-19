package engine

import (
	"context"
)

type Engine struct {
	worker map[string]*Loop
	ctx    context.Context
	stop   context.CancelFunc
}

func NewEngine() *Engine {
	ctx, stop := context.WithCancel(context.Background())

	return &Engine{
		worker: make(map[string]*Loop),
		ctx:    ctx,
		stop:   stop,
	}
}

func (e *Engine) Stop() {
	e.stop()
}

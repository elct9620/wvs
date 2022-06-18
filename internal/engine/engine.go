package engine

import (
	"context"
	"errors"
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

func (e *Engine) NewGameLoop(id string) error {
	if _, ok := e.worker[id]; ok == true {
		return errors.New("loop is created")
	}

	e.worker[id] = newLoop()
	return nil
}

func (e *Engine) Stop() {
	e.stop()
}

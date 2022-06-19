package engine

import (
	"errors"
	"sync"
	"time"
)

const FPS int = 60
const TickerDuration time.Duration = time.Millisecond / time.Duration(FPS)

type Loop struct {
	sync.Mutex
	ticker  *time.Ticker
	running bool
	exit    chan bool
}

func newLoop() *Loop {
	return &Loop{
		ticker:  time.NewTicker(TickerDuration),
		running: false,
		exit:    make(chan bool),
	}
}

func (e *Engine) NewGameLoop(id string) error {
	if _, ok := e.worker[id]; ok == true {
		return errors.New("loop is created")
	}

	e.worker[id] = newLoop()
	return nil
}

func (e *Engine) StartGameLoop(id string) error {
	loop, ok := e.worker[id]
	if ok == false {
		return errors.New("loop not exists")
	}

	if loop.running == true {
		return errors.New("loop is running")
	}

	loop.Lock()
	loop.running = true
	go func() {
		for {
			select {
			case <-loop.ticker.C:
				continue
			case <-e.ctx.Done():
				return
			case <-loop.exit:
				return
			}
		}
	}()
	loop.Unlock()

	return nil
}

func (e *Engine) StopGameLoop(id string) error {
	loop, ok := e.worker[id]
	if ok == false {
		return errors.New("loop not exists")
	}

	if loop.running == false {
		return errors.New("loop is not running")
	}

	loop.Lock()
	loop.running = false
	loop.exit <- true
	loop.Unlock()
	return nil
}

func (e *Engine) RemoveGameLoop(id string) {
	e.StopGameLoop(id)
	delete(e.worker, id)
}

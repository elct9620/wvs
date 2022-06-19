package engine

import (
	"errors"
	"sync"
	"time"
)

const FPS int = 60
const TickerDuration time.Duration = time.Millisecond / time.Duration(FPS)

type LoopFunc func(time.Duration)

type Loop struct {
	sync.Mutex
	ticker   *time.Ticker
	loopFunc LoopFunc
	running  bool
	exit     chan bool
}

func newLoop(loopFunc LoopFunc) *Loop {
	return &Loop{
		loopFunc: loopFunc,
		running:  false,
		exit:     make(chan bool),
	}
}

func (e *Engine) NewGameLoop(id string, loopFunc LoopFunc) error {
	if _, ok := e.worker[id]; ok == true {
		return errors.New("loop is created")
	}

	e.worker[id] = newLoop(loopFunc)
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
	loop.ticker = time.NewTicker(TickerDuration)
	go func() {
		previousTime := time.Now()

		for {
			select {
			case currentTime := <-loop.ticker.C:
				loop.loopFunc(currentTime.Sub(previousTime))
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

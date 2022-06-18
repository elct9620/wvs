package engine

type Engine struct {
}

func NewEngine() *Engine {
	return &Engine{}
}

func (e *Engine) NewGameLoop(id string) {
}

package db

import (
	"sync"

	"github.com/hashicorp/go-memdb"
)

type Watcher struct {
	isClosed  bool
	ch        chan *memdb.Change
	onceClose sync.Once
}

func NewWatcher() *Watcher {
	return &Watcher{
		ch: make(chan *memdb.Change),
	}
}

func (w *Watcher) Closed() bool {
	return w.isClosed
}

func (w *Watcher) Close() {
	w.onceClose.Do(func() {
		w.isClosed = true
		close(w.ch)
	})
}

package store

import (
	"errors"
	"sync"
)

type Store struct {
	sync.Mutex
	tables map[string]*Table
}

func NewStore() *Store {
	return &Store{
		tables: make(map[string]*Table),
	}
}

func (s *Store) Table(name string) *Table {
	return s.tables[name]
}

func (s *Store) CreateTable(name string) error {
	if _, ok := s.tables[name]; ok == true {
		return errors.New("table is exists")
	}

	s.Lock()
	s.tables[name] = NewTable()
	s.Unlock()

	return nil
}

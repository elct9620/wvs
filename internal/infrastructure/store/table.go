package store

import "errors"

type Table struct {
	items map[string]interface{}
}

func NewTable() *Table {
	return &Table{
		items: make(map[string]interface{}),
	}
}

func (store *Table) Find(id string) (interface{}, error) {
	if obj, ok := store.items[id]; ok == true {
		return obj, nil
	}

	return nil, errors.New("object not exists")
}

func (store *Table) Insert(id string, obj interface{}) error {
	if _, ok := store.items[id]; ok == true {
		return errors.New("object is exists")
	}

	store.items[id] = obj
	return nil
}

func (store *Table) Update(id string, obj interface{}) error {
	if _, ok := store.items[id]; ok == true {
		store.items[id] = obj
		return nil
	}

	return store.Insert(id, obj)
}

func (store *Table) Delete(id string) {
	delete(store.items, id)
}

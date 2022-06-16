package store

import "errors"

type Store struct {
	items map[string]interface{}
}

func NewStore() *Store {
	return &Store{
		items: make(map[string]interface{}),
	}
}

func (store *Store) Find(id string) (interface{}, error) {
	if obj, ok := store.items[id]; ok == true {
		return obj, nil
	}

	return nil, errors.New("object not exists")
}

func (store *Store) Insert(id string, obj interface{}) error {
	if _, ok := store.items[id]; ok == true {
		return errors.New("object is exists")
	}

	store.items[id] = obj
	return nil
}

func (store *Store) Update(id string, obj interface{}) error {
	if _, ok := store.items[id]; ok == true {
		store.items[id] = obj
		return nil
	}

	return errors.New("object not exists")
}

func (store *Store) Delete(id string) {
	delete(store.items, id)
}

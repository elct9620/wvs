package store

import "errors"

type MapFunc func(v interface{}) interface{}

type Table struct {
	items map[string]interface{}
}

func NewTable() *Table {
	return &Table{
		items: make(map[string]interface{}),
	}
}

func (table *Table) Find(id string) (interface{}, error) {
	if obj, ok := table.items[id]; ok == true {
		return obj, nil
	}

	return nil, errors.New("object not exists")
}

func (table *Table) Insert(id string, obj interface{}) error {
	if _, ok := table.items[id]; ok == true {
		return errors.New("object is exists")
	}

	table.items[id] = obj
	return nil
}

func (table *Table) Update(id string, obj interface{}) error {
	if _, ok := table.items[id]; ok == true {
		table.items[id] = obj
		return nil
	}

	return table.Insert(id, obj)
}

func (table *Table) Map(mapFunc MapFunc) []interface{} {
	items := make([]interface{}, 0)
	for _, item := range table.items {
		items = append(items, mapFunc(item))
	}

	return items
}

func (table *Table) Delete(id string) {
	delete(table.items, id)
}

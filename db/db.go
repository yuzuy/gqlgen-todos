package db

import (
	"errors"
	"sync"

	"github.com/yuzuy/gqlgen-todos/graph/model"
)

var errTodoNotFound = errors.New("todo not found")

type DB struct {
	todos map[string]*model.Todo

	mu sync.Mutex
}

func New() *DB {
	return &DB{
		todos: make(map[string]*model.Todo),
	}
}

func (d *DB) AddTodo(todo *model.Todo) error {
	d.mu.Lock()
	if _, ok := d.todos[todo.ID]; ok {
		return errors.New("id duplicated")
	}
	d.todos[todo.ID] = todo
	d.mu.Unlock()
	return nil
}

func (d *DB) FirstTodo(id string) (*model.Todo, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	todo, ok := d.todos[id]
	if !ok {
		return nil, errTodoNotFound
	}
	return todo, nil
}

func (d *DB) FindTodos() []*model.Todo {
	d.mu.Lock()
	defer d.mu.Unlock()
	todos := make([]*model.Todo, 0, len(d.todos))
	for _, v := range d.todos {
		todos = append(todos, v)
	}
	return todos
}

func (d *DB) UpdateTodo(todo *model.Todo) error {
	d.mu.Lock()
	_, ok := d.todos[todo.ID]
	if !ok {
		return errTodoNotFound
	}
	d.todos[todo.ID] = todo
	d.mu.Unlock()
	return nil
}

func (d *DB) RemoveTodo(id string) error {
	d.mu.Lock()
	if _, ok := d.todos[id]; !ok {
		return errTodoNotFound
	}
	delete(d.todos, id)
	d.mu.Unlock()
	return nil
}

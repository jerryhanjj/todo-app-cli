package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

// Todo represents a single todo item
type Todo struct {
	ID       int    `json:"id"`
	Task     string `json:"task"`
	Complete bool   `json:"complete"`
}

// TodoList represents a collection of todos
type TodoList struct {
	Todos []Todo `json:"todos"`
}

// NewTodoList creates a new empty todo list
func NewTodoList() *TodoList {
	return &TodoList{
		Todos: []Todo{},
	}
}

// Add adds a new todo to the list
func (tl *TodoList) Add(task string) {
	newID := len(tl.Todos) + 1
	todo := Todo{
		ID:       newID,
		Task:     task,
		Complete: false,
	}
	tl.Todos = append(tl.Todos, todo)
}

// List returns all todos
func (tl *TodoList) List() []Todo {
	return tl.Todos
}

// Complete marks a todo as complete
func (tl *TodoList) Complete(id int) error {
	for i := range tl.Todos {
		if tl.Todos[i].ID == id {
			tl.Todos[i].Complete = true
			return nil
		}
	}
	return fmt.Errorf("todo with ID %d not found", id)
}

// Delete removes a todo from the list
func (tl *TodoList) Delete(id int) error {
	for i := range tl.Todos {
		if tl.Todos[i].ID == id {
			tl.Todos = append(tl.Todos[:i], tl.Todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("todo with ID %d not found", id)
}

// Save saves the todo list to a file
func (tl *TodoList) Save(filename string) error {
	data, err := json.Marshal(tl)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// Load loads the todo list from a file
func (tl *TodoList) Load(filename string) error {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil // File doesn't exist, start with empty list
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, tl)
}

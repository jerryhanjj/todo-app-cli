package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jerryhanjj/todo-app-cli/internal/todo"
)

func main() {
	// Initialize todo list
	todoList := todo.NewTodoList()

	// Load existing todos
	err := todoList.Load("data/todos.json")
	if err != nil {
		fmt.Printf("Error loading todos: %v\n", err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Println("Welcome to Todo App!")
		fmt.Println("Usage: todo [add|list|complete|delete] [arguments]")
		fmt.Println("Commands:")
		fmt.Println("  add <task>        - Add a new todo")
		fmt.Println("  list              - List all todos")
		fmt.Println("  complete <id>     - Mark a todo as complete")
		fmt.Println("  delete <id>       - Delete a todo")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task to add")
			os.Exit(1)
		}
		task := os.Args[2]
		todoList.Add(task)
		fmt.Printf("Added todo: %s\n", task)

	case "list":
		todos := todoList.List()
		if len(todos) == 0 {
			fmt.Println("No todos found")
		} else {
			fmt.Println("Todos:")
			for _, t := range todos {
				status := " "
				if t.Complete {
					status = "✓"
				}
				fmt.Printf("[%s] %d: %s\n", status, t.ID, t.Task)
			}
		}

	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a todo ID to complete")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Please provide a valid todo ID")
			os.Exit(1)
		}
		err = todoList.Complete(id)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Completed todo #%d\n", id)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a todo ID to delete")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Please provide a valid todo ID")
			os.Exit(1)
		}
		err = todoList.Delete(id)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Deleted todo #%d\n", id)

	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}

	// Save todos
	err = todoList.Save("data/todos.json")
	if err != nil {
		fmt.Printf("Error saving todos: %v\n", err)
		os.Exit(1)
	}
}

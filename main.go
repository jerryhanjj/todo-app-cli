package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jerryhanjj/todo-app-cli/internal/todo"
	"github.com/spf13/pflag"
)

// getIDByIndex converts a display index (1-based) to the actual todo ID
func getIDByIndex(todos []todo.Todo, index int) (int, error) {
	if index < 1 || index > len(todos) {
		return 0, fmt.Errorf("序号 %d 超出范围 (1-%d)", index, len(todos))
	}
	return todos[index-1].ID, nil
}

func main() {
	// Initialize todo list
	todoList := todo.NewTodoList()

	// Load existing todos
	dataDir := "data"
	todoFile := filepath.Join(dataDir, "todos.json")
	err := todoList.Load(todoFile)
	if err != nil {
		fmt.Printf("Error loading todos: %v\n", err)
		os.Exit(1)
	}

	// Define command flags
	var (
		addTask    = pflag.StringP("add", "a", "", "Add a new todo")
		listTodos  = pflag.BoolP("list", "l", false, "List all todos")
		completeID = pflag.IntP("complete", "c", 0, "Mark a todo as complete by ID")
		deleteID   = pflag.IntP("delete", "d", 0, "Delete a todo by ID")
		help       = pflag.BoolP("help", "h", false, "Show help message")
	)

	// Custom usage function
	pflag.Usage = func() {
		fmt.Println("Welcome to Todo App!")
		fmt.Println("Usage: todo [OPTIONS]")
		fmt.Println("\nOptions:")
		fmt.Println("  -a, --add <task>      Add a new todo")
		fmt.Println("  -l, --list            List all todos")
		fmt.Println("  -c, --complete <id>   Mark a todo as complete")
		fmt.Println("  -d, --delete <id>     Delete a todo")
		fmt.Println("  -h, --help            Show this help message")
		fmt.Println("\nExamples:")
		fmt.Println("  todo -a \"Buy groceries\"")
		fmt.Println("  todo -l")
		fmt.Println("  todo -c 1")
		fmt.Println("  todo -d 2")
	}

	pflag.Parse()

	// Show help if requested or no flags provided
	if *help || pflag.NFlag() == 0 {
		pflag.Usage()
		os.Exit(0)
	}

	// Process add command
	if *addTask != "" {
		todoList.Add(*addTask)
		fmt.Printf("Added todo: %s\n", *addTask)
	}

	// Process list command
	if *listTodos {
		todos := todoList.List()
		if len(todos) == 0 {
			fmt.Println("No todos found")
		} else {
			fmt.Println("Todos:")
			for index, t := range todos {
				status := " "
				if t.Complete {
					status = "✓"
				}
				fmt.Printf("[%s] %d: %s\n", status, index+1, t.Task)
			}
		}
	}

	// Process complete command
	if *completeID > 0 {
		todos := todoList.List()
		actualID, err := getIDByIndex(todos, *completeID)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		err = todoList.Complete(actualID)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Completed todo #%d\n", *completeID)
	}

	// Process delete command
	if *deleteID > 0 {
		todos := todoList.List()
		actualID, err := getIDByIndex(todos, *deleteID)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		err = todoList.Delete(actualID)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Deleted todo #%d\n", *deleteID)
	}

	// Save todos
	err = todoList.Save(todoFile)
	if err != nil {
		fmt.Printf("Error saving todos: %v\n", err)
		os.Exit(1)
	}
}

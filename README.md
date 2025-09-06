# Todo App in Go

A simple command-line todo application written in Go.

## Features

- Add new todo items
- List all todo items
- Mark todos as complete
- Delete todos
- Persistent storage using JSON

## Usage

```bash
# Add a new todo
go run cmd/main.go add "Your task here"

# List all todos
go run cmd/main.go list

# Mark a todo as complete
go run cmd/main.go complete 1

# Delete a todo
go run cmd/main.go delete 1
```

## Project Structure

- `cmd/main.go` - Main application entry point
- `internal/todo/todo.go` - Todo logic and data structures
- `data/todos.json` - Persistent storage for todos

## Running the Application

1. Navigate to the project directory
2. Run commands as shown in the usage section above

The application will automatically create the data directory and todos.json file on first run.
# Todo App CLI

[![Go Version](https://img.shields.io/badge/Go-1.19+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

**Language:** [中文](README_ZH.md) | English

A simple yet powerful command-line todo application written in Go.

## ✨ Features

- 📝 Add new todo items
- 📋 List all todo items
- ✅ Mark todos as complete
- 🗑️ Delete todo items
- 💾 Persistent storage using JSON format
- 🔢 Smart index management (continuous display numbers)
- 🛡️ Secure ID generation mechanism

## 📦 Installation

### Method 1: Install using Go Install (Recommended)

**Prerequisites:** Go 1.19 or higher required

```bash
# Install the latest version
go install github.com/jerryhanjj/todo-app-cli@latest

# Or install a specific version (e.g., v1.0.0)
go install github.com/jerryhanjj/todo-app-cli@v1.0.0

# Run the application
todo-app-cli list
```

### Method 2: Install from Source

**Prerequisites:** Go 1.19 or higher required

```bash
# 1. Clone the repository
git clone https://github.com/jerryhanjj/todo-app-cli.git
cd todo-app-cli

# 2. Build the executable
go build -o todo main.go

# 3. Run the application
./todo list
```

### Method 3: Direct Run (Development Mode)

```bash
# Clone the repository
git clone https://github.com/jerryhanjj/todo-app-cli.git
cd todo-app-cli

# Run directly (recompiles each time)
go run main.go list
```

### Method 4: Using Makefile (Recommended for Developers)

```bash
# Clone the repository
git clone https://github.com/jerryhanjj/todo-app-cli.git
cd todo-app-cli

# Build using Makefile
make build

# Run the application
./todo list

# Or run directly in development mode
make run ARGS="list"
make run ARGS="add 'New Task'"
```

### Method 5: Install to System PATH

```bash
# In the project directory
go build -o todo main.go

# Move to system PATH directory (optional)
sudo mv todo /usr/local/bin/

# Or use Makefile
make install

# Now you can use it anywhere
todo list
```

## 🚀 Quick Start

```bash
# Add your first todo
./todo add "Learn Go programming"

# View all todos
./todo list

# Complete the first todo
./todo complete 1

# Delete the second todo
./todo delete 2
```

## 📖 Detailed Usage

### Adding Todos
```bash
./todo add "Buy groceries"
./todo add "Write code"
./todo add "Exercise"
```

### Viewing Todo List
```bash
./todo list
```
Example output:
```
Todos:
[✓] 1: Learn Go programming
[ ] 2: Buy groceries
[ ] 3: Write code
```

### Completing Todos
```bash
# Complete a todo using its index number
./todo complete 2
```

### Deleting Todos
```bash
# Delete a todo using its index number
./todo delete 3
```

### Getting Help
```bash
./todo
# or
./todo help
```

## 📁 Project Structure

```
todo-app-cli/
├── main.go                  # Application entry point
├── internal/
│   └── todo/
│       └── todo.go          # Todo logic and data structures
├── data/
│   └── todos.json           # JSON format persistent storage
├── Makefile                 # Build and development tools
├── go.mod                   # Go module file
├── go.sum                   # Go dependency checksum file
├── LICENSE                  # License file
├── README_ZH.md             # Project documentation (Chinese)
└── README.md                # Project documentation (English)
```

## 🔧 Development Setup

### Prerequisites
- Go 1.19+ ([Installation Guide](https://golang.org/doc/install))
- Git

### Local Development
```bash
# 1. Clone the project
git clone https://github.com/jerryhanjj/todo-app-cli.git
cd todo-app-cli

# 2. Install dependencies (if any)
go mod tidy

# 3. Use Makefile for development
make run ARGS="list"          # Run list command
make run ARGS="add 'Test'"    # Add a todo item
make build                    # Build executable
make test                     # Run tests
make fmt                      # Format code
make clean                    # Clean build files

# 4. Or use traditional way
go run main.go list
go build -o todo main.go
```

### Makefile Commands
```bash
make help          # View all available commands
make build         # Build the application
make run ARGS="..." # Run in development mode
make test          # Run tests
make clean         # Clean build files
make install       # Install to system PATH
make fmt           # Format code
make vet           # Vet code
```

## 💡 Feature Details

### Smart Index Management
- Display continuous index numbers (1, 2, 3...), even after deleting middle items
- User operations use index numbers, while internal unique IDs are maintained
- Eliminates the confusion of index gaps

### Data Persistence
- All data is saved in `data/todos.json` file
- Automatically creates data directory and file
- JSON format for easy viewing and debugging

### Error Handling
- Comprehensive boundary checking
- User-friendly error messages
- Prevention of invalid operations

## 🤝 Contributing

Contributions are welcome! Please follow these steps:

1. Fork this repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

## 🔮 Planned Features

- [ ] Add todo priority levels
- [ ] Support todo categories/tags
- [ ] Add due date functionality
- [ ] Export functionality (CSV, TXT, etc.)
- [ ] Search and filter functionality
- [ ] Colored output
- [ ] Command-line auto-completion
- [ ] Configuration file support

## 📝 Changelog

### v1.0.0 (Current Version)
- ✅ Basic CRUD functionality
- ✅ JSON persistent storage
- ✅ Smart index management
- ✅ User-friendly error messages
- ✅ Comprehensive README documentation

## 🐛 Bug Reports

If you encounter any issues or have suggestions, please:

1. [Create an Issue](https://github.com/jerryhanjj/todo-app-cli/issues)
2. Email: jerryhanjj@126.com
3. Star ⭐ this project on GitHub

## 🙏 Acknowledgments

Thanks to all developers who have contributed to this project!

Special thanks to:
- The Go community for excellent documentation and tools
- Contributors who help improve this project

---

**Happy Coding! 🚀**

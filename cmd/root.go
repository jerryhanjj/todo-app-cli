package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jerryhanjj/todo-app-cli/internal/todo"
	"github.com/spf13/cobra"
)

var (
	todoList *todo.TodoList
	todoFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "一个简单的待办事项管理工具",
	Long: `Todo App - 一个简单而强大的待办事项管理工具

使用这个工具，您可以：
- 添加新的待办事项
- 查看所有待办事项
- 标记待办事项为完成
- 删除不需要的待办事项

新的子命令方式:
  todo add "买菜"
  todo list
  todo complete 1
  todo delete 2

兼容旧版本标志方式:
  todo -a "买菜"
  todo -l
  todo -c 1
  todo -d 2`,
	Run: func(cmd *cobra.Command, args []string) {
		// 处理全局标志（兼容旧版本）
		handleGlobalFlags(cmd)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// 添加全局标志，兼容旧版本使用方式
	rootCmd.PersistentFlags().StringP("add", "a", "", "添加一个新的待办事项")
	rootCmd.PersistentFlags().BoolP("list", "l", false, "显示所有待办事项")
	rootCmd.PersistentFlags().IntP("complete", "c", 0, "标记待办事项为已完成")
	rootCmd.PersistentFlags().IntP("delete", "d", 0, "删除指定的待办事项")
}

// handleGlobalFlags processes global flags for backward compatibility
func handleGlobalFlags(cmd *cobra.Command) {
	addTask, _ := cmd.Flags().GetString("add")
	listFlag, _ := cmd.Flags().GetBool("list")
	completeID, _ := cmd.Flags().GetInt("complete")
	deleteID, _ := cmd.Flags().GetInt("delete")

	// 如果没有任何标志被设置，显示帮助
	if addTask == "" && !listFlag && completeID == 0 && deleteID == 0 {
		cmd.Help()
		return
	}

	// 处理添加命令
	if addTask != "" {
		todoList.Add(addTask)
		fmt.Printf("已添加待办事项: %s\n", addTask)
		saveTodos()
	}

	// 处理列表命令
	if listFlag {
		todos := todoList.List()
		if len(todos) == 0 {
			fmt.Println("暂无待办事项")
		} else {
			fmt.Println("待办事项:")
			for index, t := range todos {
				status := " "
				if t.Complete {
					status = "✓"
				}
				fmt.Printf("[%s] %d: %s\n", status, index+1, t.Task)
			}
		}
	}

	// 处理完成命令
	if completeID > 0 {
		todos := todoList.List()
		actualID, err := getIDByIndex(todos, completeID)
		if err != nil {
			fmt.Printf("错误: %v\n", err)
			return
		}

		err = todoList.Complete(actualID)
		if err != nil {
			fmt.Printf("错误: %v\n", err)
			return
		}
		fmt.Printf("已完成待办事项 #%d\n", completeID)
		saveTodos()
	}

	// 处理删除命令
	if deleteID > 0 {
		todos := todoList.List()
		actualID, err := getIDByIndex(todos, deleteID)
		if err != nil {
			fmt.Printf("错误: %v\n", err)
			return
		}

		err = todoList.Delete(actualID)
		if err != nil {
			fmt.Printf("错误: %v\n", err)
			return
		}
		fmt.Printf("已删除待办事项 #%d\n", deleteID)
		saveTodos()
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Initialize todo list
	todoList = todo.NewTodoList()

	// Load existing todos
	dataDir := "data"
	todoFile = filepath.Join(dataDir, "todos.json")
	err := todoList.Load(todoFile)
	if err != nil {
		fmt.Printf("Error loading todos: %v\n", err)
		os.Exit(1)
	}
}

// saveTodos saves the todo list to file
func saveTodos() {
	err := todoList.Save(todoFile)
	if err != nil {
		fmt.Printf("Error saving todos: %v\n", err)
		os.Exit(1)
	}
}

// getIDByIndex converts a display index (1-based) to the actual todo ID
func getIDByIndex(todos []todo.Todo, index int) (int, error) {
	if index < 1 || index > len(todos) {
		return 0, fmt.Errorf("序号 %d 超出范围 (1-%d)", index, len(todos))
	}
	return todos[index-1].ID, nil
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "显示所有待办事项",
	Long: `显示您的所有待办事项，包括已完成和未完成的任务。

完成的任务会显示 ✓ 标记，未完成的任务显示空格。

示例:
  todo list
  todo ls
  
  # 也可以使用标志形式（兼容旧版本）
  todo -l
  todo --list`,
	Aliases: []string{"ls", "l"},
	Run: func(cmd *cobra.Command, args []string) {
		todos := todoList.List()
		if len(todos) == 0 {
			fmt.Println("暂无待办事项")
			return
		}

		fmt.Println("待办事项:")
		for index, t := range todos {
			status := " "
			if t.Complete {
				status = "✓"
			}
			fmt.Printf("[%s] %d: %s\n", status, index+1, t.Task)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// 添加短标志支持，兼容旧版本使用方式
	listCmd.Flags().BoolP("list", "l", false, "显示所有待办事项（兼容模式）")
}

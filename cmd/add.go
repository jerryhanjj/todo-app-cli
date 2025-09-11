package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "添加一个新的待办事项",
	Long: `添加一个新的待办事项到您的待办清单中。

示例:
  todo add "买菜"
  todo add "完成项目报告"
  todo add "给妈妈打电话"
  
  # 也可以使用标志形式（兼容旧版本）
  todo -a "买菜"
  todo --add "买菜"`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var task string

		// 优先使用位置参数，如果没有则使用标志
		if len(args) > 0 {
			task = args[0]
		} else {
			task, _ = cmd.Flags().GetString("add")
		}

		if task == "" {
			fmt.Println("错误: 请提供待办事项内容")
			cmd.Help()
			return
		}

		todoList.Add(task)
		fmt.Printf("已添加待办事项: %s\n", task)
		saveTodos()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// 添加短标志支持，兼容旧版本使用方式
	addCmd.Flags().StringP("add", "a", "", "添加一个新的待办事项（兼容模式）")
}

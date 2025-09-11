package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "删除指定的待办事项",
	Long: `根据序号删除指定的待办事项。

序号对应 list 命令显示的编号。

示例:
  todo delete 1
  todo del 3
  
  # 也可以使用标志形式（兼容旧版本）
  todo -d 1
  todo --delete 1`,
	Aliases: []string{"del", "d", "rm"},
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var indexStr string

		// 优先使用位置参数，如果没有则使用标志
		if len(args) > 0 {
			indexStr = args[0]
		} else {
			flagValue, _ := cmd.Flags().GetInt("delete")
			if flagValue > 0 {
				indexStr = fmt.Sprintf("%d", flagValue)
			}
		}

		if indexStr == "" {
			fmt.Println("错误: 请提供待办事项序号")
			cmd.Help()
			return
		}

		index, err := strconv.Atoi(indexStr)
		if err != nil {
			fmt.Printf("错误: 无效的序号 '%s'\n", indexStr)
			return
		}

		todos := todoList.List()
		actualID, err := getIDByIndex(todos, index)
		if err != nil {
			fmt.Printf("错误: %v\n", err)
			return
		}

		err = todoList.Delete(actualID)
		if err != nil {
			fmt.Printf("错误: %v\n", err)
			return
		}

		fmt.Printf("已删除待办事项 #%d\n", index)
		saveTodos()
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// 添加短标志支持，兼容旧版本使用方式
	deleteCmd.Flags().IntP("delete", "d", 0, "删除指定的待办事项（兼容模式）")
}

package cmd

import (
	"aiden/expense-tracker/internal/expenses"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all expenses",
	Run: func(cmd *cobra.Command, args []string) {
		err := expenses.ListExpenses()
		if err != nil {
			fmt.Println("error listing expenses:", err)
			return
		}
	},
}

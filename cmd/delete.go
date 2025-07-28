package cmd

import (
	"aiden/expense-tracker/internal/expenses"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an expense",
	Long:  "Delete an expense",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("expense ID is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Incorrect id:", err)
			return
		}
		err = expenses.DeleteExpense(id)
		if err != nil {
			fmt.Println("Error deleting expense:", err)
		}

		fmt.Printf("Expense successfully deleted (ID: %v)\n", id)
	},
}

package cmd

import (
	"aiden/expense-tracker/internal/expenses"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&Description, "description", "d", "", "expense description")
	updateCmd.Flags().Float64VarP(&Amount, "amount", "a", 0, "expense amount")
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an expense",
	Long:  "Update an expense's description or amount",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("expense id required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if Amount < 0 {
			fmt.Println("amount cannot be negative")
			return
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}
		err = expenses.UpdateExpense(id, Description, Amount)
		if err != nil {
			fmt.Println("error updating expense:", err)
			return
		}
		fmt.Printf("Expense successfully updated (ID: %v)\n", id)
	},
}

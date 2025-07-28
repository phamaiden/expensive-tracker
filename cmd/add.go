package cmd

import (
	//"aiden/expense-tracker/internal/expenses"
	"aiden/expense-tracker/internal/expenses"
	"fmt"

	"github.com/spf13/cobra"
)

var Description string
var Amount float64

func init() {

	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&Description, "description", "d", "", "expense description (required)")
	addCmd.Flags().Float64VarP(&Amount, "amount", "a", 0, "expense amount (required)")
	addCmd.MarkFlagRequired("desc")
	addCmd.MarkFlagRequired("amt")
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an expense",
	Long:  "Add an expense with a description and amount",
	Run: func(cmd *cobra.Command, args []string) {
		if Amount < 0 {
			fmt.Println("Amount cannot be negative")
			return
		}
		id, err := expenses.AddExpense(Description, Amount)
		if err != nil {
			fmt.Println("error adding expense:", err)
			return
		}
		fmt.Printf("Expense successfully added (ID: %v)\n", id)
	},
}

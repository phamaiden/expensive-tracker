package cmd

import (
	"aiden/expense-tracker/internal/expenses"
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

var Month int
var Year int

func init() {
	rootCmd.AddCommand(summaryCmd)
	summaryCmd.Flags().IntVarP(&Month, "month", "m", 0, "integer value for month")
	summaryCmd.Flags().IntVarP(&Year, "year", "y", 0, "year")
	summaryCmd.MarkFlagsRequiredTogether("month", "year")
}

var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Prints total expenses",
	Long:  "Prints total expenses. Include month to display expenses for that month.",
	Run: func(cmd *cobra.Command, args []string) {
		month := time.Month(Month).String()
		if Month == 0 {
			if err := expenses.ExpenseSummary("all", time.Now().Year()); err != nil {
				fmt.Println("Error summarizing expenses, err")
				return
			}
			return
		}
		err := expenses.ExpenseSummary(month, Year)
		if err != nil {
			fmt.Printf("Error summarizing expenses for %s %d: %s", month, Year, err)
			return
		}
	},
}

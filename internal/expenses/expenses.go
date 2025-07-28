package expenses

import (
	"fmt"
	"time"
)

type Expense struct {
	ID          int       `json:"id"`
	Description string    `json:"desc"`
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"date"`
}

func NewExpense(id int, desc string, amt float64) *Expense {
	return &Expense{
		ID:          id,
		Description: desc,
		Amount:      amt,
		Date:        time.Now(),
	}
}

func AddExpense(desc string, amt float64) (int, error) {
	expenses, err := ReadJsonFromFile()
	if err != nil {
		return 0, err
	}

	var expenseID int
	if len(expenses) > 0 {
		expenseID = expenses[len(expenses)-1].ID
	} else {
		expenseID = 1
	}

	newExp := NewExpense(expenseID, desc, amt)
	expenses = append(expenses, *newExp)

	return expenseID, WriteJsonToFile(expenses)
}

func UpdateExpense(id int, desc string, amt float64) error {
	expenses, err := ReadJsonFromFile()
	if err != nil {
		return err
	}

	var updatedExps []Expense
	var foundExp bool

	for _, e := range expenses {
		if e.ID == id {
			foundExp = true
			if desc != "" {
				e.Description = desc
			}
			if amt >= 0 {
				e.Amount = amt
			}
		}
		updatedExps = append(updatedExps, e)
	}

	if !foundExp {
		return fmt.Errorf("Expense (ID: %v) doesn't exist", id)
	}

	return WriteJsonToFile(updatedExps)
}

func DeleteExpense(id int) error {
	expenses, err := ReadJsonFromFile()
	if err != nil {
		return err
	}

	var foundExp bool = false
	for i, exp := range expenses {
		if exp.ID == id {
			foundExp = true
			expenses = append(expenses[:i], expenses[i+1:]...)
		}
	}

	if !foundExp {
		return fmt.Errorf("Expense (ID: %v) not found", id)
	}

	return WriteJsonToFile(expenses)
}

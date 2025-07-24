package expenses

import "time"

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

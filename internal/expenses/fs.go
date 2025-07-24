package expenses

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

func GetFilePath() string {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("error getting working dir:", err)
		return ""
	}
	return path.Join(wd, "/internal/data/expenses.json")
}

func ReadJsonFromFile() ([]Expense, error) {
	filepath := GetFilePath()
	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		fmt.Println("Expense file doesn't exist, creating file")
		err := os.WriteFile(filepath, []byte("[]"), 0644)
		if err != nil {
			fmt.Println("error creating/writing to file:", err)
			return nil, err
		}

		return []Expense{}, nil
	}

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("error opening file:", err)
		return nil, err
	}

	defer file.Close()

	expenses := []Expense{}
	err = json.NewDecoder(file).Decode(&expenses)
	if err != nil {
		fmt.Println("error decoding json:", err)
		return nil, err
	}

	return expenses, nil
}

func WriteJsonToFile(expenses []Expense) error {
	filepath := GetFilePath()
	file, err := os.Create(filepath) // if file exists, truncate/clear file and write new list of tasks
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}

	err = json.NewEncoder(file).Encode(expenses)
	if err != nil {
		fmt.Println("Error encoding file:", err)
		return err
	}

	return nil
}

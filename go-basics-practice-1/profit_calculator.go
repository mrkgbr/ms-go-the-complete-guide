package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	revenue := promptForInput("Revenue: ")
	expenses := promptForInput("Expenses: ")
	taxRate := promptForInput("Tax Rate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	if err := writeResultToFile(ebt, profit, ratio); err != nil {
		log.Fatalf("Error writing results to file: %v", err)
	}

	fmt.Printf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f\n", ebt, profit, ratio)
}

func promptForInput(prompt string) float64 {
	for {
		value, err := getUserInput(prompt)
		if err != nil {
			log.Println(err) // Log the error
			continue         // Prompt again
		}
		return value
	}
}

func writeResultToFile(ebt, profit, ratio float64) error {
	result := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f\n", ebt, profit, ratio)
	return os.WriteFile("result.txt", []byte(result), 0644)
}

func calculateFinancials(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}

func getUserInput(text string) (float64, error) {
	var userInput float64
	fmt.Print(text)
	_, err := fmt.Scan(&userInput)

	if err != nil {
		return 0, errors.New("Invalid input, please enter a number.")
	}

	if userInput <= 0 {
		return 0, errors.New("Value must be a positive number.")
	}

	return userInput, nil
}

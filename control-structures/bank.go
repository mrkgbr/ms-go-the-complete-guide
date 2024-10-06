package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile, 0)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------------")
		// panic(err)
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println(randomdata.PhoneNumber())

	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your balance is: %.2f\n", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Deposit Amount: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				return
			}

			accountBalance += depositAmount
			fmt.Printf("Balance updated! Your new balance is: %.2f\n", accountBalance)
			fileops.WriteFloatToFile(accountBalanceFile, accountBalance)
		case 3:
			var withdrawalAmount float64
			fmt.Print("Withdrawal Amount: ")
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if accountBalance < withdrawalAmount {
				fmt.Println("Withdraw is aborted. Insufficient funds.")
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Printf("Balance updated! Your new balance is: %.2f\n", accountBalance)
			fileops.WriteFloatToFile(accountBalanceFile, accountBalance)
		case 4:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank!")
			return
		}
	}
}

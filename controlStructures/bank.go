package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000, errors.New("failed to fetch balance")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("failed to parse balance")
	}

	return balance, nil

}

func main() {
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----")
	}

	for i := 0; i < 200; i++ {
		fmt.Println("Welcome to Go Bank!")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Select an option: ")
		fmt.Scan(&choice)

		fmt.Println("Your choice:", choice)

		if choice == 1 {
			fmt.Printf("Your account balance is %v \n", accountBalance)
			continue
		} else if choice == 2 {
			fmt.Print("Deposit amount:")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount must be greater than 0")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Your new account balance is", accountBalance)
			writeBalanceToFile(accountBalance)
			continue
		} else if choice == 3 {
			fmt.Print("Withdraw amount:")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Print("Invalid amount must be greater than 0")
				continue
			}

			if accountBalance-withdrawAmount < 0 {
				fmt.Print("You don't have enough founds to withdraw")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("Your new account balance is", accountBalance)
			writeBalanceToFile(accountBalance)
			continue
		} else {
			fmt.Println("Goodbye!")
			break
		}
	}
}

package main

import (
	"fmt"
	"os"
)

//Goals.
//Validate user input => show error message & exit if invalid input is provided - no negative numbers - not 0.
//Store calculated results into a file.

func main() {

	revenue, revenueErr := getUserInput("Revenue: ")
	expenses, expensesErr := getUserInput("Expenses: ")
	taxRate, taxRateerr := getUserInput("Tax Rate: ")

	if expensesErr != nil || revenueErr != nil || taxRateerr != nil {
		panic("the input equals to less or equals than 0")
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
	storeResults(ebt, profit, ratio)
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT : %.1f\n Profit: %.1f\nRatio: %.3f", ebt, profit, ratio)
	os.WriteFile("profit_calculator.txt", []byte(results), 0644)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		panic("the input equals to less or equals than 0")
	}

	return userInput, nil
}

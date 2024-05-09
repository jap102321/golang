package main

import "fmt"

func main() {
	revenue := outputText("Revenue: ")
	expenses := outputText("Expenses: ")
	taxRate := outputText("Tax Rate: ")

	ebt, profit, ratio := calculateProfitValue(revenue, expenses, taxRate)

	fmt.Println("The revenue value is ", int(revenue))
	fmt.Println("The EBT is: ", ebt)
	fmt.Println(int(profit))
	fmt.Println(uint(ratio))

}

func outputText(text string) float64 {
	var userInputVariable float64
	fmt.Print(text)
	fmt.Scan(&userInputVariable)

	return userInputVariable
}

func calculateProfitValue(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	pValue := ebt * (1 - taxRate/100)
	rValue := pValue / ebt

	return ebt, pValue, rValue
}

package main

import "fmt"

func main() {

	// ----------------------- Input parameters

	revenue := getUserInput("Revenue: $ ")
	expenses := getUserInput("Expenses: $ ")
	taxRate := getUserInput("Tax Rate: $ ")

	// ----------------------- Calculation

	eBT, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	// ----------------------- Outputs

	fmt.Println("======================== Result:")
	fmt.Printf("Earning Before Tax: $ %.1f\n", eBT)
	fmt.Printf("Profit: $ %.1f\n", profit)
	fmt.Printf("Ratio: %.3f\n", ratio)
}

func getUserInput (infoText string) float64 {
	var userInput float64

	fmt.Print(infoText)
	fmt.Scan(&userInput)

	return userInput
}

func calculateFinancials (revenue, expenses, taxRate float64) (eBT float64, profit float64, ratio float64) {
	eBT = revenue - expenses
	profit = eBT * (1 - taxRate / 100)
	ratio = eBT / profit

	return
}
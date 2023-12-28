package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	// ----------------------- Input parameters

	revenue, err := getUserInput("Revenue: $ ")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	expenses, err := getUserInput("Expenses: $ ")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	taxRate, err := getUserInput("Tax Rate: $ ")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// ----------------------- Calculation

	eBT, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	// ----------------------- Outputs

	fmt.Println("======================== Result:")
	fmt.Printf("Earning Before Tax: $ %.1f\n", eBT)
	fmt.Printf("Profit: $ %.1f\n", profit)
	fmt.Printf("Ratio: %.3f\n", ratio)
	storeResults(eBT, profit, ratio)
}

func getUserInput (infoText string) (float64, error) {
	var userInput float64

	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("Value must be a positive number.")
	}

	return userInput, nil
}

func calculateFinancials (revenue, expenses, taxRate float64) (eBT float64, profit float64, ratio float64) {
	eBT = revenue - expenses
	profit = eBT * (1 - taxRate / 100)
	ratio = eBT / profit

	return
}

func storeResults(eBT, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", eBT, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}
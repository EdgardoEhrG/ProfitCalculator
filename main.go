package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	eBT := revenue - expenses
	profit := float64(eBT) * (1 - taxRate / 100)
	ratio := eBT / profit

	fmt.Println("======================== Result:")
	fmt.Println("Earning Before Tax:", eBT)
	fmt.Println("Profit:", profit)
	fmt.Println("Ratio:", ratio)
}
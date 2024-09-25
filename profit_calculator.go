package main

import (
	"fmt"
)

func profit_calculator() {
	// Build a profit calculator
	// Ask for revenue, expenses & tax rate
	// Calculate earnings before tax (EBT) and earnings after tax (profit)
	// Calculate ratio (EBT/Profit)
	// Output EBT, profit, and ratio

	var revenue, expenses, taxRate float64
	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Println(ebt, profit, ratio)

}

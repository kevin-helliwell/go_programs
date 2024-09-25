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

	revenue := requestValue("Revenue: ")
	expenses := requestValue("Expenses: ")
	taxRate := requestValue("Tax Rate: ")

	ebt, profit, ratio := generateReport(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.1f\nProfit: %.1f\nRatio: %.1f\n", ebt, profit, ratio)

}

func requestValue(text string) (value float64) {
	fmt.Print(text)
	fmt.Scan(&value)
	return value
}

func generateReport(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return ebt, profit, ratio
}

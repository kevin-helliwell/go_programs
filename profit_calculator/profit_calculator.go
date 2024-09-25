package main

import (
	"fmt"
)

// Goals
// 1) Validate user input
//    => Show error message & exit if invalid input is provided
//    - No negative numbers
//    - Not 0
// 2) Store calculated results into file

func profit_calculator() {

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

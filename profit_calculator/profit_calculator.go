package main

import (
	"errors"
	"fmt"
)

// Goals
// 1) Validate user input
//    => Show error message & exit if invalid input is provided
//    - No negative numbers
//    - Not 0
// 2) Store calculated results into file

func profit_calculator() {

	revenue, err := requestValue("Revenue: ")

	if err != nil {
		fmt.Println(err)
		return
		// panic(err)
	}

	expenses, err := requestValue("Expenses: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	taxRate, err := requestValue("Tax Rate: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	// if err1 != nil || err2 != nil || err3 != nil {
	// 	fmt.Println(err1)
	// 	return
	// }

	ebt, profit, ratio := generateReport(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.1f\nProfit: %.1f\nRatio: %.1f\n", ebt, profit, ratio)

}

func requestValue(text string) (value float64, err error) {
	fmt.Print(text)
	fmt.Scan(&value)
	if value <= 0 {
		return 0, errors.New("Value must be a positive number.")
	}
	return value, nil
}

func generateReport(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return ebt, profit, ratio
}

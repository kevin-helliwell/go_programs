package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func investment_calculator() {
	var investmentAmount float64
	expectedReturnRate := 5.5
	var years float64

	// fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	// Sprintf formats and returns a string without printing it.

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for Inflation): %.1f\n", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)

	// Printf/Println

	// fmt.Println("Future Value:", futureValue)
	// fmt.Println("Future Value (adjusted for Inflation):", futureRealValue)
	// fmt.Printf("Future Value: %.1f\nFuture Value (adjusted for Inflation): %.1f\n", futureValue, futureRealValue)

	// Multiline Strings

	// fmt.Printf(`Future Value: %.1f

	// Future Value (adjusted for Inflation): %.1f`, futureValue, futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	rfv := fv / math.Pow((1+inflationRate/100), years)
	return fv, rfv
}

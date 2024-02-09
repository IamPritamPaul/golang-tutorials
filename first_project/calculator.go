package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hiiii Pritam.")
	var investment_amount = 1000
	var expected_return = 5.5
	var years = 10
	var future_value = float64(investment_amount) * math.Pow((1+expected_return/100), float64(years))
	// var future_value = float64(investment_amount) * (1 + expected_return/100)
	fmt.Println("the invested amount is : ", future_value)
}

package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount = 1000
	var expectedReturnRate = 5.5
	var years = 10

	var futureValue = math.Pow(float64(investmentAmount)*(1+expectedReturnRate/100), float64(years))

	fmt.Println(futureValue)
}

package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 10, 15}
	sum := sumup(1, 10, 15)
	anotherSum := sumup(1, numbers...)

	fmt.Println(sum)
	fmt.Println(anotherSum)
}

func sumup(startingValue int, numbers ...int) int {
	sum := 0

	for _, v := range numbers {
		sum += v
	}
	fmt.Println("starter:", startingValue)
	return sum
}

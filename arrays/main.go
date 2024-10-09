package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5)

	userNames[0] = "Julie"
	userNames[1] = "Margo"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

	courseRatings := make(floatMap, 2)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings.output()

	for index, value := range userNames {
		fmt.Println(index, value)
	}

	for k, v := range courseRatings {
		fmt.Println(k, v)
	}
}

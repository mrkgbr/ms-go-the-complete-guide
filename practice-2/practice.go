package main

import "fmt"

type product struct {
	title string
	id    int
	price float64
}

func main() {
	// 1) Create a new array (!) that contains three hobbies you have
	// 		Output (print) that array in the command line.
	var hobbies = [3]string{"programming", "photographing", "gaming"}
	fmt.Printf("%+q\n", hobbies)
	// 2) Also output more data about that array:
	//		- The first element (standalone)
	//		- The second and third element combined as a new list
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)
	mainHobbies := hobbies[:2]

	var hobbiesFirstSlice = []string{}
	hobbiesFirstSlice = append(hobbiesFirstSlice, hobbies[0], hobbies[1])
	fmt.Println(hobbiesFirstSlice, mainHobbies)

	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)

	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	goals := []string{"learn it", "master it"}
	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	goals[1] = "find a job"
	goals = append(goals, "profit")

	fmt.Println(goals)

	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.
	products := []product{
		{"first title", 1, 4.55},
		{"second title", 2, 5.55},
	}

	products = append(products, product{"third title", 3, 6.66})

	fmt.Println(products)
}

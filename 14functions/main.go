package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")

	greeter()

	result := adder(3, 5)

	fmt.Printf("result: %v\n", result)

	proRes := proAdder(1, 2, 3, 4, 5, 6)

	fmt.Printf("proRes: %v\n", proRes)
}

func proAdder(values ...int) int {

	total := 0

	for _, v := range values {
		total += v
	}
	return total
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func greeter() {
	fmt.Println("Sille from golang")
}

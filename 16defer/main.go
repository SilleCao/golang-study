package main

import "fmt"

func main() {
	fmt.Println("Welcome to defer in golang")

	fmt.Println("Hello")
	defer fmt.Println("World1")

	fmt.Println("+++++++++++++++++++++++")

	defer fmt.Println("World2")
	fmt.Println("Hello")

	fmt.Println("+++++++++++++++++++++++")

	deferTestOne()

	fmt.Println("+++++++++++++++++++++++")

	deferTestTwo()

	fmt.Println("+++++++++++++++++++++++")

	defer deferTestFour()

	fmt.Println("+++++++++++++++++++++++")

	deferTestThree()
}

func deferTestOne() {
	fmt.Println("Hello")
	defer fmt.Println("World3")
}

func deferTestTwo() {
	defer fmt.Println("World4")
	fmt.Println("Hello")
}

func deferTestThree() {
	fmt.Println("Hello")
	defer fmt.Println("World5")
}

func deferTestFour() {
	defer fmt.Println("World6")
	fmt.Println("Hello")
}

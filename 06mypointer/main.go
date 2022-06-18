package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	var ptr *int

	fmt.Println("Value of pointer is ", ptr)

	myNum := 23

	var nptr = &myNum

	fmt.Println("Value of acutal pointer is ", nptr)
	fmt.Println("Value of acutal pointer is ", *nptr)

	*nptr = *nptr * 2
	fmt.Println("New value is ", myNum)

}

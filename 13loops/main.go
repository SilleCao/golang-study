package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Firday", "Saturday"}

	fmt.Printf("days: %v\n", days)

	fmt.Println("===============================")
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	fmt.Println("===============================")
	for i := range days {
		fmt.Println(days[i])
	}

	fmt.Println("===============================")
	for i, v := range days {
		fmt.Printf("index %v , value is %v\n", i, v)
	}

	fmt.Println("===============================")

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 5 {
			break
		}

		fmt.Println("Value is ", rougueValue)
		rougueValue++

	}

	fmt.Println("===============================")

	rougueValue2 := 1

	for rougueValue2 < 10 {

		if rougueValue2 == 4 {
			rougueValue2++
			continue
		}

		fmt.Println("Value is ", rougueValue2)
		rougueValue2++

	}

	fmt.Println("===============================")

	rougueValue3 := 1

	for rougueValue3 < 10 {

		if rougueValue3 == 2 {
			goto lco
		}

		fmt.Println("Value is ", rougueValue3)
		rougueValue3++

	}

lco:
	fmt.Println("Hello World!")

}

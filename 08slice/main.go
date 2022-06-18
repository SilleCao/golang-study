package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitList: %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Printf("fruitList: %v\n", fruitList)

	fruitList = append(fruitList[1:])
	fmt.Printf("fruitList: %v\n", fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Printf("fruitList: %v\n", fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 777
	fmt.Printf("highScores: %v\n", highScores)

	highScores = append(highScores, 555, 666, 321)

	fmt.Printf("highScores: %v\n", highScores)

	fmt.Printf("sort.IntsAreSorted(highScores): %v\n", sort.IntsAreSorted(highScores))

	sort.Ints(highScores)
	fmt.Printf("highScores: %v\n", highScores)

	//how to remove a value from slices based on index
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Printf("courses: %v\n", courses)
	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Printf("courses: %v\n", courses)
}

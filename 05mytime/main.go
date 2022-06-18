package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()

	fmt.Printf("presentTime: %v\n", presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	fmt.Println(presentTime.Format("MM-dd-yyyy HH:mm:ss")) //error format

	t := time.Date(2022, time.June, 18, 22, 18, 0, 0, time.Now().Location())

	fmt.Println(t)
	fmt.Println(t.Format("2006-01-02 15:04:05"))
}

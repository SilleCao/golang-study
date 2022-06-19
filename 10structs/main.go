package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	//no inheritance in golang; No super or parent

	sille := User{"Sille", "caoliang1025@163.com", true, 18}
	fmt.Println(sille)

	fmt.Printf("Sille details are: %+v\n", sille)
	fmt.Printf("Name is %v and email is %v\n", sille.Name, sille.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	//no inheritance in golang; No super or parent

	sille := User{"Sille", "caoliang1025@163.com", true, 18}
	fmt.Println(sille)

	fmt.Printf("Sille details are: %+v\n", sille)
	fmt.Printf("Name is %v and email is %v\n", sille.Name, sille.Email)

	sille.GetStatus()
	sille.GetStatusWithoutPointer()

	sille.NewEmail("sille@go.dev")
	fmt.Printf("sille: %v\n", sille)

	sille.NewEmailWithoutPinter("sille@go.dev1")
	fmt.Printf("sille without pointer: %v\n", sille)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u *User) GetStatus() {

	fmt.Println("Is user active: ", u.Status)

}

func (u User) GetStatusWithoutPointer() {

	fmt.Println("Is user active without using pointer: ", u.Status)

}

func (u *User) NewEmail(email string) {

	u.Email = email

}

func (u User) NewEmailWithoutPinter(email string) {

	u.Email = email

}

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Printf("languages: %v\n", languages)
	fmt.Printf("languages[\"JS\"]: %v\n", languages["JS"])

	delete(languages, "RB")
	fmt.Printf("languages: %v\n", languages)

	//loops are interesting in golang
	for k, v := range languages {
		fmt.Printf("for key %v, value is %v\n", k, v)
	}

	for _, v := range languages {
		fmt.Printf("key v. value is %v\n", v)

	}
}

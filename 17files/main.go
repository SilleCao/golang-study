package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")

	content := "What are you going up there?"

	file, err := os.Create("./mytestfile.txt")

	if err != nil {
		panic(err)
	}

	len, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Printf("len: %v\n", len)

	defer file.Close()

	readFile("./mytestfile.txt")

}

func readFile(filename string) string {

	databyte, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	fmt.Println("Text data inside the file is ", databyte)
	fmt.Println("===============================================")
	s := string(databyte)
	fmt.Println("Text data inside the file is ", s)
	return s

}

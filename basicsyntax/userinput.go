package main

import "fmt"

func main() {
	var fName string
	fmt.Println("Please enter First Name: ")
	fmt.Scanln(&fName)

	var lName string
	fmt.Println("Enter Last Name: ")
	fmt.Scanln(&lName)

	fmt.Println("Name ", &lName)
}

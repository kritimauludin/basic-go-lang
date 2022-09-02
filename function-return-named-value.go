package main

import "fmt"

func getData() (firstName string, lastName string, age int8) {
	firstName = "kriti"
	lastName = "mauludin"
	age = 20

	return
}

func main() {
	firstName, lastName, age := getData()

	fmt.Println(firstName, lastName)
	fmt.Println(age)
}
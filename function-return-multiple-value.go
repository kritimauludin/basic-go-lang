package main

import "fmt"

func getFullName() (string, string, int8) {
	return "kriti", "mauludin", 20
}

func main() {
	// firstName, lastName, age := getFullName()
	// fmt.Println(firstName)
	// fmt.Println(lastName)
	// fmt.Println(age)

	//abaikan atau ignore beberapa return
	firstName, _, _ := getFullName()
	fmt.Println(firstName)
}
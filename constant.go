package main

import "fmt"

func main() {
	const firstName string = "kriti"
	const lastname = "mauludin"

	// multiple constant
	const (
		footLength = 43
		age = 3
	)

	fmt.Println(firstName)

	fmt.Println(footLength)
	fmt.Println(age)
}
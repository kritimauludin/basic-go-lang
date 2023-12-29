package main

import "fmt"

func main() {
	var name string
	var home = "cilendek"
	var age int8 = 20
	//hanya untuk deklarasi awal tanpa var
	footLength := 43

	name = "kriti mauludin"
	fmt.Println(name)
	fmt.Println(home)
	fmt.Println(age)

	fmt.Println(footLength)
	footLength = 42
	fmt.Println(footLength)

	name = "kriti"
	fmt.Println(name)

	var (
		firstName = "kriti"
		lastName  = "mauludin"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

}

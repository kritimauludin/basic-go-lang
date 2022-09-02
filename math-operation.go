package main

import "fmt"

func main() {
	var result = 10 + 5
	fmt.Println(result)

	var a= 10
	var b = 10
	result = a + b
	fmt.Println(result)

	// augmenter asigment
	var i = 10
	i+=1
	fmt.Println(i)

	// unary operator
	i++
	fmt.Println(i)

	var negative = -100
	fmt.Println(negative)
}
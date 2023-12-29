package main

import "fmt"

// func sayHelloWithFilter(name string, filter func(string) string) {
// 	nameFiltered := filter(name)

// 	fmt.Println("hello ", nameFiltered)
// }

// with type declaration jika penyebutan function akan sangat panjang
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)

	fmt.Println("hello ", nameFiltered)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("kriti", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("anjing", filter)
}

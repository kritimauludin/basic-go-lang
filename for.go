package main

import (
	"fmt"
)

func main() {
	//for dengan statement
	for counter:=1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}

	// for dengan range cocok untuk loop aray, slice, dan map
	//manual
	slice := []string{"kriti", "mauludin", "mauludin kriti"}
	for i:=0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	//for range
	for index, value := range slice {
		fmt.Println("index", index, "=", value)
	}
	// jika key/index tidak dibutuhkan bisa gunakan _
	for _, value := range slice {
		fmt.Println(value)
	}
}
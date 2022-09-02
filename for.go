package main

import (
	"fmt"
)

func main() {
	// counter := 1

	// digo hanya ada 1 perulangan yaitu for loops
	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke ", counter)
	// 	counter++
	// }

	//for dengan statement
	// init statement hanya dijalankan diawal dan hanya sekali dieksekusi
	// post statement akan dijalankan setelah blok code perulang
	for counter:=1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}

	// for dengan range cocok untuk loop aaray slice map

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

	person := make(map[string]string)
	person ["name"] = "kriti" 
	person ["jobs"] = "mahasiswa"
		
	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
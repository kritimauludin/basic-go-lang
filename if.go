package main

import "fmt"

func main() {
	var name = "kriti"

	if name == "kriti" {
		fmt.Println("Hello kriti")
	} else if name == "mauludin" {
		fmt.Println("Hello mauludin")
	} else {
		fmt.Println("siapa anda?")
	}

	//long statements
	// var length = len(name);

	// if length > 5 {
	// 	fmt.Println("nama terlalu panjang")
	// }else{
	// 	fmt.Println("nama sudah benar")
	// }

	// short statements
	if length := len(name); length > 5 {
		fmt.Println("nama terlalu panjang")
	}else{
		fmt.Println("nama sudah benar")
	}
}
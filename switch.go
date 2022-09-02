package main

import "fmt"

func main() {
	name := "mauludin mauludin"

	switch name {
	case "kriti":
		fmt.Println("halo kriti")
	case "mauludin" :
		fmt.Println("halo mauludin")
		fmt.Println("namanya panjang")
	default :
		fmt.Println("siapa kamu?")
	}

	//short statement switch
	// switch length := len(name); length > 5{
	// case true :
	// 	fmt.Println("nama kepanjangan")
	// case false :
	// 	fmt.Println("nama sudah benar")
	// }

	// switch tanpa kondisi variabel lgsg masuk dicase
	// ini mirip dengan if statement
	// sehingga penggunaan switch digo sangat flexible
	length := len(name)

	switch {
	case length > 10 : 
		fmt.Println("nama kepanjangan")
	case length > 5 :
		fmt.Println("nama lumayan panjang")
	default : 
		fmt.Println("nama sudah benar")
	}
}
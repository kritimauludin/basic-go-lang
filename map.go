package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "kriti",
		"address": "bogor",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	person["jobs"] = "mahasiswa"
	fmt.Println(person)

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go"
	book["author"] = "Kang Eko"
	book["wrong"] = "delete"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "wrong")

	fmt.Println(book)
	fmt.Println(len(book))
}
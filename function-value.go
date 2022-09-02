package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {

	//untungnya nanti ketika funtion dibutuhkan parameter function lain
	//function di golang adalah first class citizen (tidak dianaktirikan)
	sayGoodBye := getGoodBye

	fmt.Println(sayGoodBye("kriti"))
	fmt.Println(getGoodBye("mauludin"))

}
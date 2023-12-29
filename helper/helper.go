package helper

import "fmt"

//ini private tidak bs diakses dr luar
var version = 1.0

// ini public atau bisa diakses dr luar
var Application = "Go Dasar"

//bisa diakses dr luar package
func SayHello(name string){
	fmt.Println("hello ", name)
}

func sayGoodbye(name string){
	fmt.Println("Goodbye", name)
}
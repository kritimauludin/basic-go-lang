package main

import "fmt"

func sayHelloTo(name string, age int8) {
	fmt.Println("hello", name, "your age", age)

}

func main(){
	sayHelloTo("kriti", 20)
}
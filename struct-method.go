package main

import "fmt"

type Employee struct {
	Name, Job string
	Age       int8
}

/** 
	sebenernya misalkan kita buat function spesifik hanya untuk struct Employee
	diusahan buat function dalam bentuk struct method. agar seakan akan struct
	employee memiliki sebuah method
*/
func (employee Employee) sayHello(name string) {
	fmt.Println("Hello", name, "My name is", employee.Name)
}

func main(){
	kriti := Employee{
		Name : "Kriti",
		Job : "Web Dev",
		Age : 20,
	}

	kriti.sayHello("mauludin")
}
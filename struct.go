package main

import "fmt"

//programmer go bisasanya menggunakan uppercase to lower dlm membuat struct

type Employee struct {
	Name, Position string
	Age            int8
}

func main() {
	//cara buat go - 1
	var kriti Employee
	kriti.Name ="Kriti"
	kriti.Position = "Web Developer"
	kriti.Age = 20

	fmt.Println(kriti)
	fmt.Println(kriti.Position)

	//cara buat go - 2
	mauludin := Employee{
		Name: "Mauludin",
		Position: "Backend Developer",
		Age: 20,
	}
	fmt.Println(mauludin)

	//cara buat go - 3
	//pastikan posisinya sesuai, namun jika struct diubah akan error
	//maka lebih baik gunakan cara 2 / 1 dibanding cara ini karena rentan error
	andi := Employee{"andi", "System Analyst", 25}
	fmt.Println(andi)

}
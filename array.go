package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "kriti"
	names[1] = "mauludin"
	names[2] = "mauludin kriti"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		80,
		70,
	}
	fmt.Println(values)
	fmt.Println(values[0])

	fmt.Println(len(names))
	fmt.Println(len(values))

	var again = [...]int{ //ketika ingin declare seperti ini, harus lgsg diisi
		90,
		70,
		80,
		98,
		0,
	}
	fmt.Println(len(again))
}

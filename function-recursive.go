package main

import "fmt"

//pastikan harus ada kondisi berhenti ketika menggunakan recursive

//cara 1 dengan for loop
func factorialForLoop(value int) int{
	result :=1

	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

//cara 2 dengan recursif
func factorialRecursive(value int) int{
	if value == 1 {
		return 1
	} else{
		return value * factorialRecursive(value- 1)
	}
}

func main(){
	loop := factorialForLoop(8)
	recursive := factorialRecursive(8)

	fmt.Println(loop)
	fmt.Println(recursive)
}
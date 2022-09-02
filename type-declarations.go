package main

import "fmt"

func main() {
	type NoKtp string
	type Work bool

	var ktpKm NoKtp = "221121212122"
	var workStatus Work = false
	fmt.Println(ktpKm)
	fmt.Println(workStatus)

}
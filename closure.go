package main

import "fmt"

//hati hati ketika membuat function didalam function
//hindari merubah value data/ variabel yang tidak diperlukan

//scope yg ada diatas bisa diakses dibawah
//tetapi scope yang ada dibawah tidak bisa diakses diatasnya

//clousure merupakan kemampuan suatu function mengakses data data disekitarnya

func main() {
	counter := 0
	name := "km"

	increment := func() {
		fmt.Println("increment")
		counter++

		//jika tidak ingin merubah variabel diatasnya
		//diperlukan pendeklarasian ulang
		name := "kriti"
		fmt.Println(name)
	}

	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)

}

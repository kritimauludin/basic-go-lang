package main

import "fmt"

func Lah(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "lah"
	}
}

func main() {
	//kita tidak bisa langsung menggunakan interface kosong dimasukan ke variabel
	//maka kita harus set terlebih dahulu variabelnya menjadi interface kosong. contoh
	var data interface{} = Lah(1)
	data2 := Lah(2)
	data3	:= Lah(3)

	fmt.Println(data)
	fmt.Println(data2)
	fmt.Println(data3)
}
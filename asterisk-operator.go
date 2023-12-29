/**
Operator *

● Saat kita mengubah variable pointer, maka yang berubah hanya variable tersebut.
● Semua variable yang mengacu ke data yang sama tidak akan berubah
● Jika kita ingin mengubah seluruh variable yang mengacu ke data tersebut, kita bisa menggunakan
operator *
*/

package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Bogor", "Jawa Barat", "Indonesia"}
	address2 := &address1

	address2.City = "Cipanas"

	fmt.Println(address1)
	fmt.Println(address2)

	// address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
}

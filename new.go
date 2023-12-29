/**
Operator new

● Sebelumnya untuk membuat pointer dengan menggunakan operator &
● Go-Lang juga memiliki function new yang bisa digunakan untuk membuat pointer
● Namun function new hanya mengembalikan pointer ke data kosong, artinya tidak ada data awal
*/

package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 *Address = new(Address)
	var address2 *Address = address1

	address2.City = "Cipanas"

	fmt.Println(address1)
	fmt.Println(address2)
}

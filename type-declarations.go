/*
*
Type Declarations adalah kemampuan membuat ulang tipe data baru dari tipe data yang sudah ada
Type Declarations biasanya digunakan untuk membuat alias terhadap tipe data yang sudah ada,
dengan tujuan agar lebih mudah dimengerti
*/
package main

import "fmt"

func main() {
	type NoKtp string
	type Work bool

	var ktpKm NoKtp = "221121212122"
	var workStatus Work = false
	fmt.Println(ktpKm)
	fmt.Println(NoKtp("111111111")) //konversi no ktp jadi string
	fmt.Println(workStatus)

}

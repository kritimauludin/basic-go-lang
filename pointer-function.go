/**
tips jika kita memili struct yg fieldnya banyak misal 50 field, lebih baik menggunakan pointer pada functionnya
karena jika tidak, maka setiap pemanggilan parameter. struct tersebut akan terus diduplikat
akibatnya jika parameter digunakan dlm banyak function maka memori akan semakin membengkak
*/

package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// yang dikirim pada parameter disini adalah data duplikat dari variabel (pass by value)
func ChangeCountryToIndonesia(address Address) {
	address.Country = "Indonesia"
}

//tanda bintang memberi tahu bahwa Addressnya haruslah pointer jgn duplikatnya
func ChangeCountryToIndonesiaWithPointer(address *Address){
	address.Country = "Indonesia"
}

func main(){
	var address = Address{"Bogor", "Jawa Barat", ""}

	ChangeCountryToIndonesia(address)

	fmt.Println(address) //country tidak berubah karena data tidak mereference

	//pemanggilan function ditambah & karena sebelumnya variabel address adalah data aslinya
	//sehingga ketika ditambah & tipenya menjadi pointer
	ChangeCountryToIndonesiaWithPointer(&address)
	fmt.Println(address) //berubah
}
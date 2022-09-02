/**
secara default di go variabel itu passing by value bukan reference
jadi jika mengirim varibel ke dlm func, method atau variabel lain
sebenernya yg dikirim adalah duplikasi valuenya
*/

/**
pointer adalah kemampuan membuat reference ke lokasi data dimemory yang sama
tanpa menduplikasi data yang sudah ada

sederhanya, dengan pointer. kita bisa membuat pass by reference
*/

package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main(){
	address1 := Address{"Bogor", "Jawa Barat", "Indonesia"}
	address2 := address1

	address2.City = "Cipanas"

	fmt.Println(address1)
	fmt.Println(address2)


	//implement pass by reference dengan operator and(&)
	address3 := &address1 //pointer

	address3.City = "Bandung"

	//ini menyebabkan addres 3 pindah pointer/ membuat memory baru
	// address3 = &Address{"solo", "jawa timur", "Indonesia"} 

	//jika kita ingin merubah value dari siapapun yg mengacu kesuatu memori
	//maka bisa gunakan operator * seperti dibawah ini
	*address3 = Address{"solo", "jawa timur", "Indonesia"} 



	fmt.Println(address1)
	fmt.Println(address3)


/** 
sebelumnya untuk buat pointer dengan menggunakan operator &
go juga memiliki function new yg bisa digunakan utk membuat pointer
namun function new hanya mengembalikan pointer kedata kosong artinya tdk ada data awal
*/

var addres4 *Address = new(Address)

addres4.City = "Bandung"
fmt.Println(addres4)



}
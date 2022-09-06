/**
* Pada Array atau Slice, untuk mengakses data, kita menggunakan index Number dimulai dari 0
* Map adalah tipe data lain yang berisikan kumpulan data yang sama, namun kita bisa menentukan jenis tipe data index yang akan kita gunakan
* Sederhananya, Map adalah tipe data kumpulan key-value (kata kunci - nilai), dimana kata kuncinya bersifat unik, tidak boleh sama
* Berbeda dengan Array dan Slice, jumlah data yang kita masukkan ke dalam Map boleh sebanyak-banyaknya, asalkan kata kunci nya berbeda, 
  jika kita gunakan kata kunci sama, maka secara otomatis data sebelumnya akan diganti dengan data baru
*/
package main

import "fmt"

func main() {

	//buat map
	person := map[string]string{
		"name":    "kriti",
		"address": "bogor",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	person["jobs"] = "mahasiswa"
	fmt.Println(person)

	//membuat map baru
	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go"
	book["author"] = "Kang Eko"
	book["wrong"] = "delete"
	fmt.Println(book)
	fmt.Println(len(book))

	//dlete value map with key
	delete(book, "wrong")

	fmt.Println(book)
	fmt.Println(len(book))
}

package main

// Interface Kosong

// ● Go-Lang bukanlah bahasa pemrograman yang berorientasi objek
// ● Biasanya dalam pemrograman berorientasi objek, ada satu data parent di puncak yang bisa
// dianggap sebagai semua implementasi data yang ada di bahasa pemrograman tersebut
// ● Contoh di Java ada java.lang.Object
// ● Untuk menangani kasus seperti ini, di Go-Lang kita bisa menggunakan interface kosong
// ● Interface kosong adalah interface yang tidak memiliki deklarasi method satupun, hal ini membuat
// secara otomatis semua tipe data akan menjadi implementasi nya
// ● Interface kosong, juga memiliki type alias bernama any
// ● Ada banyak contoh penggunaan interface kosong di Go-Lang, seperti :
// ○ fmt.Println(a ...interface{})

import "fmt"

// func Lah(i int) interface{} {
func Lah(i int) any {
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
	data3 := Lah(3)

	fmt.Println(data)
	fmt.Println(data2)
	fmt.Println(data3)
}

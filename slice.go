/*
*
Tipe data Slice adalah potongan dari data Array
Slice mirip dengan Array, yang membedakan adalah ukuran Slice bisa berubah
Slide dan Array selalu terkoneksi, dimana Slice adalah data yang mengakses sebagian atau seluruhdata di Array

Tipe Data Slice memiliki 3 data, yaitu pointer, length dan capacity
Pointer adalah penunjuk data pertama di array para slice
Length adalah panjang dari slice, dan
Capacity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capacity
*/
package main

import "fmt"

func main() {
	var months = [12]string{
		"jan",
		"feb",
		"mar",
		"apr",
		"mei",
		"jun",
		"jul",
		"aug",
		"sep",
		"okt",
		"nov",
		"des",
	}

	// slice reference from array
	var slice1 = months[4:7]
	fmt.Println(slice1)
	// get length
	fmt.Println(len(slice1))
	// get capacity
	fmt.Println(cap(slice1))

	months[5] = "juni"
	fmt.Println(slice1)

	slice1[2] = "Juli"
	fmt.Println(slice1)

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "bulan baru")
	fmt.Println(slice3)

	slice3[1] = "Bukan desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	//buat slice baru
	newSlice := make([]string, 2, 4)
	newSlice[0] = "kriti"
	newSlice[1] = "mauludin"
	// newSlice[2] = "nama"  erorr harusnya menggunakan append
	fmt.Println(newSlice)

	newSlice2 := append(newSlice, "namaku")
	fmt.Println(newSlice2)

	//copy slice (make sure length and capacity same, so that the data is not truncate)
	fromSlice := months[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// code aray vs slice
	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}

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

	fmt.Println(newSlice)

	//copy slice (make sure length and capacity same, so that the data is not truncate)
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	// code aray vs slice
	iniArray:= [5]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
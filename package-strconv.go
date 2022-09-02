/**
sebelumnya kita sudah belajar cara konversi tipe data, misal dari int32 ke int34

untuk menkonversi misal tipe data yg berbeda dari int ke string atau sebaliknya
kita bisa lakukan dengan bantuan package strconv(string conversion)

https://golang.go/pkg/strconv/

*/

package main

import (
	"fmt"
	"strconv"
)

func main(){
	//kektika ingin merubah dr string ke tipe data yg dituju gunakan(parse)
	//ketika ingin merubah dr tipe data tertentu ke string gunakan (format)

	// untuk melakukan konversi string ke bool func parsebool akan mengembalikan 2 nilai
	// value dan errornya
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)	
	} else {
		fmt.Println("Errot" , err.Error())
	}

	number, err := strconv.ParseInt("100000", 10, 32)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println("Errot" , err.Error())
	}

	// argument kedua yaitu base 2 8 10 16
	value := strconv.FormatInt(100000, 2)
	fmt.Println(value)

	valueInt, err := strconv.Atoi("200000")
	if err == nil {
		fmt.Println(valueInt)
	}
}
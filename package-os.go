/**
Go-Lang telah menyediakan banyak sekali package bawaan, salah satunya package os

package os berisikan fungsionalitas untuk mengakses fitur sistem operasi secara
independen (bisa digunakan disemua sistem operasi)

https://golang.go/pkg/os/

*/

package main

import (
	"fmt"
	"os"
)

func main(){
	args := os.Args
	fmt.Println("Argument : ")
	fmt.Println(args)
	// cara ambil data argument misal : go run os.go kriti mauludin
	// fmt.Println(args[1])

	// hotsname akan mengembalikan 2 nilai, hosname itu sendiri dan error
	hostname, err := os.Hostname()

	if err == nil {
		fmt.Println(hostname)
	}else {
		fmt.Println("Error : ", err.Error())
	}

	//contoh untuk mendapatkan environment variabel, 
	// biasanya digunakan untuk mensetting konfigurasi aplikasi
	// misal ingin digunakan untk koneksi ke database
	GOPATH := os.Getenv("GOPATH")

	fmt.Println(GOPATH)
}
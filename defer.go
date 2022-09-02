package main

import "fmt"

//defer adalah function yang akan dipanggil jika suatu function selesai dijalankan
//defer tidak memperdulikan function itu error atau tidak

func logging(){
	fmt.Println("selamat logging dieksekusi with ")
}

func runApplication(value int){
	//biasakan gunakan defer diawal agar tidak bertemu error
	defer logging()

	fmt.Println("run application dieksekusi")
	result := 10/value
	fmt.Println("Result", result)
}

func main()  {
	runApplication(10)
	runApplication(0)
}
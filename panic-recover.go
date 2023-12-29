package main

import "fmt"

func endApp() {
	fmt.Println("End App")
	//recover berfungsi menangkap data panic agar aplikasi tidak dihentikan
	//pastikan recover dieksekusi diakhir maka kita bisa simpan di defer
	//dalam artian recover yg benar disimpan didefer
	message := recover() //cara yg bener recover disimpan di function defeer
	if message != nil {
		fmt.Println("Error with message: ", message)
	}

}

func runApp(error bool) {
	defer endApp()
	if error {

		//jika ada sistem eeror dan ingin aplikasi berhenti maka gunakan panic
		panic("Error App")
	}
	//jika melakukan recover disini maka tidak akan berjalan
	//kita hanya bisa menangkap data panic diluar dari function ini

	fmt.Println("Running App")
}

func main() {
	runApp(false)
}

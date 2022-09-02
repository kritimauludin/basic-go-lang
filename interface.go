//jangan keliru interfase disini bukan UI interface

/**
jadi interface disini merupakan tipe data abstract yang tidak diimplmentasi langsung,
jadi murni sebagai kontrak yang berisi definisi-definisi method
*/

package main

import "fmt"

//interface seperti kontrak baru
//dan biasanya interface ini digunakan di function-function yang general
//
type HasName interface{
	//misal dikontrak ini kita ingin dia memiliki function get name dengan return value string
	GetName() string
}

func sayHalo(hasName HasName){
	fmt.Println("Halo", hasName.GetName())
}

//contoh penggunaan interface dengan struct - 1
type Person struct{
	Name string
}

//karena struct person memiliki function get name, 
//maka secara otomatis struct person ini sudah mengikuti/memenuhi kontrak interface HasName
func (person Person) GetName() string{
	return person.Name
}

//contoh penggunaan interface dengan struct - 2

//walaupun structnya berbeda kita bisa menggunakan interface yang sama. Contoh
type Animal struct{
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main(){
	//cara memasukan datanya tidak bisa seperti struct
	//sehingga cara mengimplementasikan datanya seperti sebagai berikut
	//data tersebut harus memiliki function GetName
	kriti := Person{
		Name: "Kriti",
	}

	sayHalo(kriti)

	//cara gunakan interface - 2
	var gajah Animal
	gajah.Name = "Gajah"
	sayHalo(gajah)
}
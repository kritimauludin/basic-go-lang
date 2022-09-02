/**
biarpun method akan menempel di struct, tapi sebenarnya data struct yang diakses
dalam method adalah pass by value.

sangat direkomendasikan menggunakan pointer di method
sehingga tidak boros memori karena harus selalu diduplikasi ketika memanggil method

*/

package main

import "fmt"

//contoh func/method bukan gunakan pointer
type Man struct{
	Name string
}

func (man Man) Married() {
	man.Name = "Mr. " + man.Name
	fmt.Println("Di Method", man.Name)
}

//contoh func/method gunakan pointer
type Girl struct{
	Name string
}

// dengan pointer makan akan hemat memori karena tidak menduplikasi data yg dijadikan parameter
func (girl *Girl) Married() {
	girl.Name = "Mrs. " + girl.Name
}



func main()  {
	//penggunaan tanpa pointer
	agus := Man{"Agus"}
	agus.Married()

	fmt.Println(agus.Name)

	//penggunaan dengan pointer
	dina := Girl{"Dina"}
	dina.Married()

	fmt.Println(dina.Name)
}
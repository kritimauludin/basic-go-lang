package main

import "fmt"

/**
type asertion adlh kemampuan merubah tipe data ke tipe data yg diinginkan
fitur ini sering kali digunakan ketika bertemu dengan data interface kosong
*/

//pastikan tipe data interface kosong sesuai dengan tipe data yg diinginkan
//jika interface kosong mengembalikan string dan kita merubah ke int maka akan terjadi panic

func random() interface{} {
	return "Test"
}

func main(){
	result := random()
	resultString := result.(string)
	fmt.Println(resultString)

	//jika ingin lebih aman bisa menggunakan switch agar jika terjadi panic bisa ter-recover
	resultRandom := random()
	switch value := resultRandom.(type){
	case string :
		fmt.Println("String", value)
	case int :
		fmt.Println("Int", value)
	default :
		fmt.Println("unknown")
	}

}
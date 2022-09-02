/**
dalam bahasa pemograman biasanya ada fitur reflection, dimana kita bisa melihat
struktur kode kita pada saat aplikasi sedang berjalan

hal ini bisa dilakukan digo-lang dengan menggunakan package reflect

package ini perlu dipelajari secara lebih lanjut dengan otodidak

reflection sangat berguna ketika kita ingin membuat library yang general sehingga mudah digunakan

https://golang.go/pkg/reflect/
*/
package main

import (
	"fmt"
	"reflect"
)

type Sample struct{
	//membuat tag sangat cocok jika kita ingin membuat library untuk melakukan validasi
	//dibandingkan cek manual dengan if else
	Name string `required:"true" max:"10"`
}

type ContohLagi struct {
	Name string `required:"true"`
	Desc string `required:"true"`
}

//contoh kasus kita akan membuat validation library terhadap sebuah data struuct
func IsValid(data interface{}) bool{
	//kita ambil data reflectionnya
	t := reflect.TypeOf(data)
	//lakukan looping
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {

			//logika pendeknya
			// return reflect.ValueOf(data).Field(i).Interface() != ""
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}

func main(){
	sample := Sample{"Kriti"}

	//untuk membuat reflectionnya
	var sampleType reflect.Type = reflect.TypeOf(sample)
	structField := sampleType.Field(0)

	fmt.Println(sampleType.NumField())
	fmt.Println(structField.Name)

	//untuk mendapatkan structagnya bisa menggunakan
	required := structField.Tag.Get("required")
	max := structField.Tag.Get("max")
	min := structField.Tag.Get("min")

	fmt.Println(required)
	fmt.Println(max)
	fmt.Println(min)

	fmt.Println(IsValid(sample))

	contoh := ContohLagi{"kriti", "orang"}
	contoh2:= ContohLagi{"", "orang"}
	fmt.Println(IsValid(contoh))
	fmt.Println(IsValid(contoh2))
}
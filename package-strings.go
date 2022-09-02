/**

ini adalah package yg berisikan function-function untuk memanipulasi tipe data string
ada banyak sekali function yang bisa kita gunakan

https://golang.org/pkg/strings
*/

package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println(strings.Contains("kriti mauludin", "agus"))
	fmt.Println(strings.Contains("kriti mauludin", "kriti"))

	//menghapus karakter dikanan dan kiri dimana disini separatornya dengan spasi
	//salah satu penggunakannya untuk memparsing input user 
	fmt.Println(strings.Trim("  Kriti Mauludin   ", " "))

	//memotong string menjadi slice
	fmt.Println(strings.Split("kriti mauludin", " "))

	fmt.Println(strings.ToLower("Kriti Mauludin"))
	fmt.Println(strings.ToUpper("Kriti Mauludin"))

	fmt.Println(strings.ReplaceAll("makan malam ketika makan malam", "malam", "siang"))

}
package main

//saat kita lakukan import secra otomatis func init di connection akan dijalankan
import (
	"fmt"
	"go-dasar/database"
)

func main(){
	result := database.GetDatabase()

	fmt.Println(result)
}
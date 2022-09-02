/**
package initialization

saat buat package, kita bs membuat sbuah func yg akan diakses ketika package kita diakses/diimport

kita cukup membuat func dengan nama init disebuah package file agar func init tersebut scra lgsg dieksekusi

ini sangat cocok ketika misal, jika package berisi function-function untuk berkomunikasi dengan
database, kita membuat function inisialisasi untuk membuka koneksi ke database.
*/

package database

import "fmt"

var connection string

func init() {
	fmt.Println("init dijalankan")
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
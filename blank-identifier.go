/**
kadang kita hanya ingin menjalankan init function dipackage tanpa harus mengekskusi
salah satu function yang ada didalam package

secara default, G0-Lang akan komplen ketika ada package yang diimport namum tidak digunakan

untuk menangani hal tersebut, kita bisa menggunakan blank identifier(_) sebelum nAMA
package ketika melakukan import
*/

package main

import _ "go-dasar/database"

func main(){
	
}
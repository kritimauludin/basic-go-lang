/**
package container/ring meupakan adalah implementasi struktur data circular list

circular list(ring/lingkaran) adalah struktur data ring, dimana diakhir element akan kembali ke element awal(HEAD)

https://golang.go/pkg/container/ring/
*/

package main

import (
	"container/ring"
	"fmt"
	"strconv"
)


func main(){
	//tentukan jumlah ringnya
	var data *ring.Ring = ring.New(5)
	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	//function bawaan ring untuk mengiterate datanya
	data.Do(func(value interface{}) {
		fmt.Println(value)
	})

	//jangan gunakan for loop karena akan terjadi infinit loop karena data ring tidak ada akhirnya
}

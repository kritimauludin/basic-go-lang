/**
package time adalah package yang berisikan fungsionaliras untuk management waktu di Go-Lang

jadi semua ada reprenstasinya dengan nama time
https://golang.go/pkg/time/
*/

package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println(time.Now())

	//untuk membuat customisasi waktu secara manual
	utc := time.Date(2022, time.June, 10, 22, 0, 0, 0, time.UTC)
	fmt.Println(utc.Local())

	//misal kita ingin parsing tanggal yg didapat dr api dalam bentuk string

	//adapun layout digo sedikit unik karena tidak bisa dengan "YYYY-DD-MM" seperti diibahasa lain seperti php dan java
	//digo untuk layout menggunakan format dalam tanggal dan tahun
	// layout := "time.RFC3339"
	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "2020-01-02")
	fmt.Println(parse)


}

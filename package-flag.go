/**
package flag berisikan fungsionalitas untuk memparsing command line argumen

ini sangat cocok jika kita ingin buat aplikasi berbasis terminal karna kita tidak
perlu repot repot membuat helpernya, melakukan parsing datanya karena semua sudah
dihandle oleh package flag

https://golang.go/pkg/flag/
*/

package main

import (
	"flag"
	"fmt"
)

func main(){
	// .string digunakan agar nantinya diparsing dlm bentuk string dan value parsingnya akn disimpan divariabel
	var host *string = flag.String("host", "localhost", "Put your database host")
	username := flag.String("username", "root", "Put your database username")
	password := flag.String("password", "root", "Put your database password")
	var port *int = flag.Int("port", 8080, "Put yout database port")
	
	//perlu melakukan proses parsing, sebelum menggunakan variabel-variabel sebelumnya
	flag.Parse()


	/**
	hasil varibel-varibel diatas adalah bentuk pointer, sehingga ketika kita ingin mengambil value perlu tanda (*)
	*/
	fmt.Println("host : ", *host)
	fmt.Println("username : ", *username)
	fmt.Println("password : ", *password)
	fmt.Println("port : ", *port)

	//cara jalankan : go run flag.go -host={} -username={} -password={}
}
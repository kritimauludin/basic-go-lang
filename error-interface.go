/**
digolang kita tidak perlu membuat struct error secara manual
karena golang sudah menyediakan library untuk membuat helper secara mudah
library ini terdapat di package errors
*/
package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai int, pembagi int) (int, error){
	if pembagi == 0 {
		return 0, errors.New("Pembagian Dengan NOL")
	}else {
		//karena error adlh interface maka kita bisa set dengan nil jika tdk butuh
		result := nilai/pembagi
		return result, nil
	}
}

func main(){
	hasil, err := pembagian(10000000000, 5000)

	if err == nil {
		fmt.Println("Hasil : ", hasil)
	}else {
		fmt.Println("Error : ", err.Error())
	}

}
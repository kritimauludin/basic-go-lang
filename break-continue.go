/**
* Break & continue adalah kata kunci yang bisa digunakan dalam perulangan
* Break digunakan untuk menghentikan seluruh perulangan
* Continue adalah digunakan untuk menghentikan perulangan yang berjalan, dan langsungmelanjutkan ke perulangan selanjutnya
*/
package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("perulangan ke ", i)
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println("ganjil = ", i)
	}
}

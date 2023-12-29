package main

// Parameter yang berada di posisi terakhir, memiliki kemampuan dijadikan sebuah varargs
// ● Varargs artinya datanya bisa menerima lebih dari satu input, atau anggap saja semacam Array.
// ● Apa bedanya dengan parameter biasa dengan tipe data Array?
// ○ Jika parameter tipe Array, kita wajib membuat array terlebih dahulu sebelum mengirimkan ke function
// ○ JIka parameter menggunakan varargs, kita bisa langsung mengirim data nya, jika lebih dari satu, cukup
// gunakan tanda koma

import "fmt"

// jika ingin ada 2 parameter dalam variadic function
// pastikan variabel argumen berada dipaling akhir
func sumAll(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}

	return total
}

func main() {
	total := sumAll(10, 11, 23, 32, 33)
	total2 := sumAll()

	fmt.Println(total)
	fmt.Println(total2)

	slice := []int{20, 21, 9, 10, 22}

	//... untuk menandakan bahwa itu merupakan slice
	total3 := sumAll(slice...)
	fmt.Println(total3)
}

package main

/**
tidak semua tipe data bisa menggunakan nil karena terkadang ada yg lgsg kena default value
contoh : int menjadi 0, string menjadi string kosong dsb
nil hanya bisa digunakan pada interface, function, map, slice, pointer dan channel
*/

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		/** 
			nil diperlukan untuk melakukan pengecekan
			jadi semisal ada pengecekan apakah data nil/tidak 
			kita tidak perlu masuk ke dalam mapnya untuk mengambil key yang ingin dicek
		*/
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {

	//contoh pengecekan tanpa nil tanpa nil
	var result map[string]string
	if result["name"] == ""{
		fmt.Println("data kosong")
	}else {
		fmt.Println(result)
	}

	//contoh pengecekan dengan nil akan lebih simple
	result2 := NewMap("ahmad")
	if result2 == nil{
		fmt.Println("data kosong")
	}else {
		fmt.Println(result2)
	}

}
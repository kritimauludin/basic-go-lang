/**
package sort meupakan pakcage yg berisikan utilitas(istilahnya helper) untuk proses pengurutan

agar data kita bisa diurutkan, kita harus mengimplementasikan kontrak diinterface sort.Interface

https://golang.go/pkg/sort/
*/

package main

import (
	"fmt"
	"sort"
)

type User struct{
	Name string
	Age int
}

//kita buat type alias dengan nama UserSlice []User
//dimana ini adalah array of user
type UserSlice []User


// untuk mengurutkan kita perlu kontrak Interface yang terdiri Len, Less, Swap
func (userSlice UserSlice) Len() int {
	return len(userSlice)
}

//dalam hal ini misal kita ingin mengurutkannyta berdasarkan Age maka contohnya seperti in
func (userSlice UserSlice) Less(i, j int) bool{
	return userSlice[i].Age > userSlice[j].Age

	//jika ingin dengan nama kita tinggal mencomparenya dengan Name
	// return userSlice[i].Name < userSlice[j].Name
}

func (userSlice UserSlice) Swap(i, j int){
	//kita bisa swap sepeti biasa dengan temporary variabel
	//namun digo ada cara cepatnya seperti dibawah ini
	userSlice[i], userSlice[j] = userSlice[j], userSlice[i]
}

func main(){
	users := []User{
		{"kriti", 19},
		{"mauludin", 20},
		{"adit", 25},
		{"ahmad", 10},
	}

	fmt.Println(users)
	sort.Sort(UserSlice(users))
	fmt.Println(users)
}
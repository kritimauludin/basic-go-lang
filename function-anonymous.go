//function tanpa nama
package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked ", name)
	} else {
		fmt.Println("welcome ", name)
	}
}

func main(){

	//cara 1
	blacklist := func (name string) bool {
		return name == "admin"
	}

	registerUser("adi", blacklist)
	registerUser("admin", blacklist)

	// cara 2
	registerUser("root", func (name string) bool {
		return name == "root"
	})
}
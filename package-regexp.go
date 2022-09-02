/**
package refgexp (regular expression)
package regexp adalah utilitas digo-lang untuk melakukan pencarian regular expression

regular expression digo-lang menggunakan library c yang dibuat google bernama RE2

ini penting sekali untk dipelajari karena kedepannya ketika membuat validasi, ataupun bikin
apapun yg berhubungan dengan text maka regula expression ini akan sangat membantu

https://github.com/google/re2/wiki/Syntax
https://golang.go/pkg/regexp/
*/

package main

import (
	"fmt"
	"regexp"
)

func main(){
	//untuk membuat object regexp kita harus menggunakan function MustCompile kemudian
	//masukan pola/pattern regluar expressionnya

	//membuat regexp
	var regek = regexp.MustCompile(`([a-z]).com`)
	//mengecek apakah regexp match dengan string
	fmt.Println(regek.MatchString("s.com"))
	fmt.Println(regek.MatchString("s.id"))
	fmt.Println(regek.MatchString("a.com"))

	// mencari string yang match dengan maximum jumlah hasil
	search := regek.FindAllString("s.com b.com b.com b.id", 1)
	search2 := regek.FindAllString("s.com b.com b.com b.id", -1)
	fmt.Println(search)
	fmt.Println(search2)
}
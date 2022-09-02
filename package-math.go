/**
package math meupakan package yg berisi constant dan fungsi matematika

https://golang.go/pkg/math/
*/

package main

import (
	"fmt"
	"math"
)

func main(){
	//membulakan bilangan float keatas/kebawah
	fmt.Println(math.Round(1.7))
	fmt.Println(math.Round(1.3))

	//memaksa membulatkan bilangan float kebawah
	fmt.Println(math.Floor(1.7))

	//memaksa membulatkan bilangan float keatas
	fmt.Println(math.Ceil(1.3))

	fmt.Println(math.Max(17, 20))
	fmt.Println(math.Min(10, 20))

	
}
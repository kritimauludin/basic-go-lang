/**
package container/list meupakan adalah implementasi struktur data double linked list di Go-Lang

https://golang.go/pkg/container/list/
*/

package main

import (
	"container/list"
	"fmt"
)

func main(){
	data := list.New()
	data.PushBack("kriti")
	data.PushBack("mauludin")
	data.PushBack("agus")
	data.PushFront(2)

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Next())

	//dari depan kebelakang
	for e := data.Front(); e != nil; e = e.Next(){
		fmt.Println(e.Value)
	}

	//dari belakang kedepan
	for element := data.Front(); element != nil; element = element.Next(){
		fmt.Println(element.Value)
	}
}
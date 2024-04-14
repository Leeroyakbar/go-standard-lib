package main

import (
	list2 "container/list"
	"fmt"
)

func main() {
	list := new(list2.List)
	list.PushBack("Lee")
	list.PushBack("Roy")
	list.PushBack("Akbar")
	//fmt.Println(list.Front().Value)
	//fmt.Println(list.Front().Next().Value)

	// atau
	for e := list.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

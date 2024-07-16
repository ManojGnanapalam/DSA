package chapter01

import (
	"container/list"
	"fmt"
)

func ListDSA() {
	var list list.List

	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(4)
	list.PushBack(5)

	for i := list.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	fmt.Printf("The type %T", list)
}

package main

import (
	"fmt"
	"generic-collections/list"
)

func main() {
	data := []int{1, 2, 3, 4, 5}
	numList := list.From(data...)
	fmt.Println(numList)
	fmt.Println(numList.Contains(7))
}

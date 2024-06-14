package main

import (
	"fmt"
	"generic-collections/list"
)

func main() {
	data := []int{1, 2, 3, 4, 5}
	numList := list.From(data...)

	var arr = numList.ToSlice()
	arr[0] = 100

	fmt.Println(&numList.ToSlice()[0], &arr[0])

	for i := 0; i < numList.Count(); i++ {
		//var element = numList.Get(i)
		//fmt.Println(&element, &arr[i], &element == &arr[i])
		fmt.Println(&arr[i])
	}
}

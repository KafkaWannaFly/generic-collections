package list

import (
	"fmt"
	"generic-collections/interfaces"
	"generic-collections/utils"
)

type List[T any] struct {
	elements []T
	count    int
}

func New[T any]() *List[T] {
	return &List[T]{elements: make([]T, 0)}
}

func From[T any](elements ...T) *List[T] {
	var list = &List[T]{elements: elements}
	list.count = len(elements)
	return list
}

func (receiver *List[T]) ForEach(appliedFunc func(int, T)) {
	for i, element := range receiver.elements {
		appliedFunc(i, element)
	}
}

func (receiver *List[T]) Add(item T) interfaces.ICollection[T] {
	receiver.elements = append(receiver.elements, item)
	receiver.count++

	return receiver
}

func (receiver *List[T]) AddAll(items interfaces.ICollection[T]) interfaces.ICollection[T] {
	receiver.elements = append(receiver.elements, items.ToSlice()...)
	receiver.count = len(receiver.elements)
	return receiver
}

func (receiver *List[T]) Count() int {
	return receiver.count
}

func (receiver *List[T]) Contains(item T) bool {
	for _, element := range receiver.elements {
		if utils.IsEqual(element, item) {
			return true
		}
	}

	return false
}

func (receiver *List[T]) ContainsAll(items interfaces.ICollection[T]) bool {
	elementMap := make(map[string]bool)
	receiver.ForEach(func(index int, element T) {
		elementMap[fmt.Sprintf("%v", element)] = true
	})

	var result = true
	items.ForEach(func(index int, item T) {
		if _, exists := elementMap[fmt.Sprintf("%v", item)]; !exists {
			result = false
		}
	})

	return result
}

func (receiver *List[T]) Clear() interfaces.ICollection[T] {
	receiver.elements = make([]T, 0)
	receiver.count = 0

	return receiver
}

func (receiver *List[T]) Filter(predicate func(T) bool) interfaces.ICollection[T] {
	ans := New[T]()

	for _, element := range receiver.elements {
		if predicate(element) {
			ans.Add(element)
		}
	}

	return ans
}

func (receiver *List[T]) Get(i any) T {
	var index = i.(int)
	return receiver.elements[index]
}

func (receiver *List[T]) Set(i any, item T) interfaces.ICollection[T] {
	var index = i.(int)
	receiver.elements[index] = item

	return receiver
}

func (receiver *List[T]) ToSlice() []T {
	return receiver.elements
}

func (receiver *List[T]) IsEmpty() bool {
	return receiver.count == 0
}

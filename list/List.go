package list

import (
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

func (receiver *List[T]) Has(item T) bool {
	for _, element := range receiver.elements {
		if utils.IsEqual(element, item) {
			return true
		}
	}

	return false
}

func (receiver *List[T]) HasAll(items interfaces.ICollection[T]) bool {
	elementMap := make(map[string]bool)
	receiver.ForEach(func(index int, element T) {
		var key = utils.HashCodeOf(element)
		elementMap[key] = true
	})

	var result = true
	items.ForEach(func(index int, item T) {
		var key = utils.HashCodeOf(item)
		if _, exists := elementMap[key]; !exists {
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

func (receiver *List[T]) Get(i int) T {
	return receiver.elements[i]
}

func (receiver *List[T]) Set(i int, item T) {
	receiver.elements[i] = item
}

func (receiver *List[T]) ToSlice() []T {
	return receiver.elements
}

func (receiver *List[T]) IsEmpty() bool {
	return receiver.count == 0
}

func (receiver *List[T]) Clone() interfaces.ICollection[T] {
	return From[T](receiver.elements...)
}

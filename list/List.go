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

func (receiver *List[T]) Add(item T) {
	receiver.elements = append(receiver.elements, item)
	receiver.count++
}

func (receiver *List[T]) AddAll(items interfaces.ICollection[T]) {
	receiver.elements = append(receiver.elements, items.ToSlice()...)
	receiver.count = len(receiver.elements)
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
	var result = true

	for _, element := range items.ToSlice() {
		if !receiver.Contains(element) {
			result = false
			break
		}
	}

	return result
}

func (receiver *List[T]) Clear() {
	receiver.elements = make([]T, 0)
	receiver.count = 0
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

func (receiver *List[T]) Set(i any, item T) {
	var index = i.(int)
	receiver.elements[index] = item
}

func (receiver *List[T]) ToSlice() []T {
	return receiver.elements
}

func (receiver *List[T]) IsEmpty() bool {
	return receiver.count == 0
}

func (receiver *List[T]) Remove(item T) {
	for i, element := range receiver.elements {
		if utils.IsEqual(element, item) {
			receiver.elements = append(receiver.elements[:i], receiver.elements[i+1:]...)
			receiver.count--
			break
		}
	}
}

func (receiver *List[T]) RemoveAll(items interfaces.ICollection[T]) {
	var notRemoveIndices = make(map[int]bool, 0)

	for _, element := range receiver.elements {
		items.ForEach(func(index int, item T) {
			if !utils.IsEqual(element, item) {
				notRemoveIndices[index] = true
			}
		})
	}

	newSize := receiver.Count() - items.Count()
	newCollection := make([]T, 0)
	for k, _ := range notRemoveIndices {
		newCollection = append(newCollection, receiver.elements[k])
	}

	receiver.elements = newCollection
	receiver.count = newSize
}

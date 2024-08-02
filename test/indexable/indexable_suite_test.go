package indexable_test

import (
	"generic-collections/interfaces"
	"generic-collections/linkedlist"
	"generic-collections/list"
	"generic-collections/utils"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type _IIndexable[TIndex any, TValue any] interface {
	interfaces.IIndexableGetSet[TIndex, TValue]
	interfaces.IIndexableAdder[TIndex, TValue]
	interfaces.IIndexableRemover[TIndex, TValue]
	interfaces.IIndexableFinder[TIndex, TValue]
}

type _IIndexableCollection[TIndex any, TValue any] interface {
	interfaces.ICollection[TValue]
	_IIndexable[TIndex, TValue]
}

func TestIndexable(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Indexable Suite")
}

var _ = Describe("Test IndexableCollection", func() {

	When("Using integer", func() {
		integerLinkedList := linkedlist.From(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		Context("For LinkedList", integerTests(integerLinkedList))

		integerList := list.From(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		Context("For List", integerTests(integerList))
	})

	When("Using string", func() {
		stringLinkedList := linkedlist.From("a", "b", "c", "d", "e", "f", "g", "h", "i", "j")
		Context("For LinkedList", stringTests(stringLinkedList))

		stringList := list.From("a", "b", "c", "d", "e", "f", "g", "h", "i", "j")
		Context("For List", stringTests(stringList))
	})
})

func integerTests(collection interfaces.ICollection[int]) func() {
	return func() {
		var integerCollection _IIndexableCollection[int, int]

		BeforeEach(func() {
			integerCollection = collection.Clone().(any).(_IIndexableCollection[int, int])

			Expect(integerCollection.Count()).To(Equal(10))
			Expect(collection.ToSlice()).To(Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
		})

		It("Should be able to get element by index", func() {
			for i := 0; i < integerCollection.Count(); i++ {
				Expect(integerCollection.GetAt(i)).To(Equal(i + 1))
			}
		})

		It("Should panic when get element out of range", func() {
			Expect(func() { integerCollection.GetAt(999999999999999) }).To(Panic())
		})

		It("Should try get element in range", func() {
			val, ok := integerCollection.TryGetAt(0)
			Expect(val).To(Equal(1))
			Expect(ok).To(BeTrue())
		})

		It("Should try get element out of range", func() {
			val, ok := integerCollection.TryGetAt(-99999999999999)
			Expect(val).To(Equal(utils.DefaultValue[int]()))
			Expect(ok).To(BeFalse())
		})

		It("Should be able to set element by index", func() {
			integerCollection.SetAt(0, 100)
			Expect(integerCollection.GetAt(0)).To(Equal(100))
		})

		It("Should panic when set element out of range", func() {
			Expect(func() { integerCollection.SetAt(999999999999999, 100) }).To(Panic())
		})

		It("Should try set element in range", func() {
			ok := integerCollection.TrySetAt(0, 100)
			Expect(ok).To(BeTrue())
			Expect(integerCollection.GetAt(0)).To(Equal(100))
		})

		It("Should try set element out of range", func() {
			ok := integerCollection.TrySetAt(-99999999999999, 100)
			Expect(ok).To(BeFalse())
		})

		It("Should add element to the beginning", func() {
			integerCollection.AddFirst(999)

			Expect(integerCollection.Count()).To(Equal(11))
			Expect(integerCollection.GetAt(0)).To(Equal(999))
		})

		It("Should add element to the end", func() {
			integerCollection.AddLast(999)

			Expect(integerCollection.Count()).To(Equal(11))
			Expect(integerCollection.GetAt(integerCollection.Count() - 1)).To(Equal(999))
		})

		It("Should add element before a certain index", func() {
			integerCollection.AddBefore(5, 999)

			Expect(integerCollection.Count()).To(Equal(11))
			Expect(integerCollection.GetAt(5)).To(Equal(999))

			integerCollection.AddBefore(0, 888)

			Expect(integerCollection.Count()).To(Equal(12))
			Expect(integerCollection.GetAt(0)).To(Equal(888))

			integerCollection.AddBefore(integerCollection.Count(), 777)

			Expect(integerCollection.Count()).To(Equal(13))
			Expect(integerCollection.GetAt(integerCollection.Count() - 1)).To(Equal(777))
		})

		It("Should be panic when add element before a certain index but out of range", func() {
			Expect(func() { integerCollection.AddBefore(999999999999999, 999) }).To(Panic())

			Expect(func() { integerCollection.AddBefore(-999999999999999, 999) }).To(Panic())
		})

		It("Should try to add element before a certain index", func() {
			ok := integerCollection.TryAddBefore(5, 999)
			Expect(ok).To(BeTrue())
			Expect(integerCollection.GetAt(5)).To(Equal(999))

			ok = integerCollection.TryAddBefore(0, 888)
			Expect(ok).To(BeTrue())
			Expect(integerCollection.GetAt(0)).To(Equal(888))

			ok = integerCollection.TryAddBefore(integerCollection.Count(), 777)
			Expect(ok).To(BeTrue())
			Expect(integerCollection.GetAt(integerCollection.Count() - 1)).To(Equal(777))
		})

		It("Should try to add element before a certain index but out of range", func() {
			ok := integerCollection.TryAddBefore(999999999999999, 999)
			Expect(ok).To(BeFalse())

			ok = integerCollection.TryAddBefore(-999999999999999, 999)
			Expect(ok).To(BeFalse())
		})

		It("Should add element after a certain index", func() {
			integerCollection.AddAfter(5, 999)

			Expect(integerCollection.Count()).To(Equal(11))
			Expect(integerCollection.GetAt(6)).To(Equal(999))

			integerCollection.AddAfter(0, 888)

			Expect(integerCollection.Count()).To(Equal(12))
			Expect(integerCollection.GetAt(1)).To(Equal(888))

			integerCollection.AddAfter(integerCollection.Count()-1, 777)

			Expect(integerCollection.Count()).To(Equal(13))
			Expect(integerCollection.GetAt(integerCollection.Count() - 1)).To(Equal(777))

			integerCollection.AddAfter(-1, 666)

			Expect(integerCollection.Count()).To(Equal(14))
			Expect(integerCollection.GetAt(0)).To(Equal(666))
		})

		It("Should try add element after a certain index", func() {
			ok := integerCollection.TryAddAfter(5, 999)
			Expect(ok).To(BeTrue())
			Expect(integerCollection.GetAt(6)).To(Equal(999))

			ok = integerCollection.TryAddAfter(0, 888)
			Expect(ok).To(BeTrue())
			Expect(integerCollection.GetAt(1)).To(Equal(888))

			ok = integerCollection.TryAddAfter(integerCollection.Count()-1, 777)
			Expect(ok).To(BeTrue())
			Expect(integerCollection.GetAt(integerCollection.Count() - 1)).To(Equal(777))

			ok = integerCollection.TryAddAfter(-1, 666)
			Expect(ok).To(BeTrue())
			Expect(integerCollection.GetAt(0)).To(Equal(666))
		})

		It("Should try add element after a certain index but out of range", func() {
			ok := integerCollection.TryAddAfter(999999999999999, 999)
			Expect(ok).To(BeFalse())

			ok = integerCollection.TryAddAfter(-999999999999999, 999)
			Expect(ok).To(BeFalse())
		})

		It("Should remove first element", func() {
			first := integerCollection.RemoveFirst()
			Expect(integerCollection.Count()).To(Equal(9))
			Expect(first).To(Equal(1))
			Expect(integerCollection.GetAt(0)).To(Equal(2))

			first = integerCollection.RemoveFirst()
			Expect(integerCollection.Count()).To(Equal(8))
			Expect(first).To(Equal(2))
			Expect(integerCollection.GetAt(0)).To(Equal(3))

			first = integerCollection.RemoveFirst()
			Expect(integerCollection.Count()).To(Equal(7))
			Expect(first).To(Equal(3))
			Expect(integerCollection.GetAt(0)).To(Equal(4))
		})

		It("Should remove last element", func() {
			last := integerCollection.RemoveLast()
			Expect(integerCollection.Count()).To(Equal(9))
			Expect(last).To(Equal(10))
			Expect(integerCollection.GetAt(integerCollection.Count() - 1)).To(Equal(9))

			last = integerCollection.RemoveLast()
			Expect(integerCollection.Count()).To(Equal(8))
			Expect(last).To(Equal(9))
			Expect(integerCollection.GetAt(integerCollection.Count() - 1)).To(Equal(8))

			last = integerCollection.RemoveLast()
			Expect(integerCollection.Count()).To(Equal(7))
			Expect(last).To(Equal(8))
			Expect(integerCollection.GetAt(integerCollection.Count() - 1)).To(Equal(7))
		})

		It("Should remove element by index", func() {
			integerCollection.RemoveAt(0)
			Expect(integerCollection.Count()).To(Equal(9))
			Expect(integerCollection.GetAt(0)).To(Equal(2))

			integerCollection.RemoveAt(integerCollection.Count() - 1)
			Expect(integerCollection.Count()).To(Equal(8))
			Expect(integerCollection.GetAt(integerCollection.Count() - 1)).To(Equal(9))
		})

		It("Should panic when remove element out of range", func() {
			Expect(func() { integerCollection.RemoveAt(999999999999999) }).To(Panic())

			Expect(func() { integerCollection.RemoveAt(-999999999999999) }).To(Panic())
		})

		It("Should try remove element by index", func() {
			val, ok := integerCollection.TryRemoveAt(0)
			Expect(integerCollection.Count()).To(Equal(9))
			Expect(val).To(Equal(1))
			Expect(ok).To(BeTrue())
		})

		It("Should try remove element by index but out of range", func() {
			val, ok := integerCollection.TryRemoveAt(999999999999999)
			Expect(val).To(Equal(utils.DefaultValue[int]()))
			Expect(ok).To(BeFalse())

			val, ok = integerCollection.TryRemoveAt(-999999999999999)
			Expect(val).To(Equal(utils.DefaultValue[int]()))
			Expect(ok).To(BeFalse())
		})

		It("Should find the first element based on predicate", func() {
			index := integerCollection.FindFirst(func(i int, val int) bool { return val == 5 })
			Expect(index).To(Equal(4))

			index = integerCollection.FindFirst(func(i int, val int) bool { return val == 999 })
			Expect(index).To(Equal(-1))
		})

		It("Should find the last element based on predicate", func() {
			index := integerCollection.FindLast(func(i int, val int) bool { return val == 5 })
			Expect(index).To(Equal(4))

			index = integerCollection.FindLast(func(i int, val int) bool { return val == 999 })
			Expect(index).To(Equal(-1))
		})

		It("Should find all elements based on predicate", func() {
			indexes := integerCollection.FindAll(func(i int, val int) bool { return val > 5 })
			Expect(indexes).To(Equal([]int{5, 6, 7, 8, 9}))

			indexes = integerCollection.FindAll(func(i int, val int) bool { return val < 0 })
			Expect(indexes).To(Equal([]int{}))
		})
	}
}

func stringTests(collection interfaces.ICollection[string]) func() {
	return func() {
		var stringCollection _IIndexableCollection[int, string]

		BeforeEach(func() {
			stringCollection = collection.Clone().(any).(_IIndexableCollection[int, string])

			Expect(stringCollection.Count()).To(Equal(10))
		})

		It("Should add an element at the head", func() {
			stringCollection.AddFirst("z")

			Expect(stringCollection.Count()).To(Equal(11))
			Expect(stringCollection.GetAt(0)).To(Equal("z"))
		})

		It("Should add an element at the tail", func() {
			stringCollection.AddLast("z")

			Expect(stringCollection.Count()).To(Equal(11))
			Expect(stringCollection.GetAt(10)).To(Equal("z"))
		})

		It("Should add before an element", func() {
			stringCollection.AddBefore(0, "head")
			stringCollection.AddBefore(stringCollection.Count(), "tail")

			Expect(stringCollection.Count()).To(Equal(12))
			Expect(stringCollection.GetAt(0)).To(Equal("head"))
			Expect(stringCollection.GetAt(stringCollection.Count() - 1)).To(Equal("tail"))

			stringCollection.AddBefore(5, "middle")

			Expect(stringCollection.Count()).To(Equal(13))
			Expect(stringCollection.GetAt(5)).To(Equal("middle"))
		})

		It("Should add after an element", func() {
			stringCollection.AddAfter(-1, "head")
			stringCollection.AddAfter(stringCollection.Count()-1, "tail")

			Expect(stringCollection.Count()).To(Equal(12))
			Expect(stringCollection.GetAt(0)).To(Equal("head"))
			Expect(stringCollection.GetAt(stringCollection.Count() - 1)).To(Equal("tail"))

			stringCollection.AddAfter(5, "middle")

			Expect(stringCollection.Count()).To(Equal(13))
			Expect(stringCollection.GetAt(6)).To(Equal("middle"))
		})

		It("Should try add before an element in range", func() {
			ok := stringCollection.TryAddBefore(0, "head")
			Expect(ok).To(BeTrue())
			Expect(stringCollection.Count()).To(Equal(11))
			Expect(stringCollection.GetAt(0)).To(Equal("head"))
		})

		It("Should try add before an element out of range", func() {
			ok := stringCollection.TryAddBefore(-1, "head")
			Expect(ok).To(BeFalse())
			Expect(stringCollection.Count()).To(Equal(10))
		})

		It("Should try add after an element in range", func() {
			ok := stringCollection.TryAddAfter(-1, "head")
			Expect(ok).To(BeTrue())
			Expect(stringCollection.Count()).To(Equal(11))
			Expect(stringCollection.GetAt(0)).To(Equal("head"))
		})

		It("Should try add after an element out of range", func() {
			ok := stringCollection.TryAddAfter(stringCollection.Count(), "tail")
			Expect(ok).To(BeFalse())
			Expect(stringCollection.Count()).To(Equal(10))
		})

		It("Should remove first item", func() {
			stringCollection.RemoveFirst()

			Expect(stringCollection.Count()).To(Equal(9))
			Expect(stringCollection.GetAt(0)).To(Equal("b"))
		})

		It("Should remove last item", func() {
			stringCollection.RemoveLast()

			Expect(stringCollection.Count()).To(Equal(9))
			Expect(stringCollection.GetAt(8)).To(Equal("i"))
		})

		It("Should get an element", func() {
			Expect(stringCollection.GetAt(0)).To(Equal("a"))
			Expect(stringCollection.GetAt(1)).To(Equal("b"))
			Expect(stringCollection.GetAt(2)).To(Equal("c"))
			Expect(stringCollection.GetAt(3)).To(Equal("d"))
			Expect(stringCollection.GetAt(4)).To(Equal("e"))
			Expect(stringCollection.GetAt(5)).To(Equal("f"))
			Expect(stringCollection.GetAt(6)).To(Equal("g"))
			Expect(stringCollection.GetAt(7)).To(Equal("h"))
			Expect(stringCollection.GetAt(8)).To(Equal("i"))
			Expect(stringCollection.GetAt(9)).To(Equal("j"))
		})

		It("Should set an element", func() {
			stringCollection.SetAt(0, "aa")
			stringCollection.SetAt(1, "bb")
			stringCollection.SetAt(2, "cc")
			stringCollection.SetAt(3, "dd")
			stringCollection.SetAt(4, "ee")
			stringCollection.SetAt(5, "ff")
			stringCollection.SetAt(6, "gg")
			stringCollection.SetAt(7, "hh")
			stringCollection.SetAt(8, "ii")
			stringCollection.SetAt(9, "jj")

			Expect(stringCollection.GetAt(0)).To(Equal("aa"))
			Expect(stringCollection.GetAt(1)).To(Equal("bb"))
			Expect(stringCollection.GetAt(2)).To(Equal("cc"))
			Expect(stringCollection.GetAt(3)).To(Equal("dd"))
			Expect(stringCollection.GetAt(4)).To(Equal("ee"))
			Expect(stringCollection.GetAt(5)).To(Equal("ff"))
			Expect(stringCollection.GetAt(6)).To(Equal("gg"))
			Expect(stringCollection.GetAt(7)).To(Equal("hh"))
			Expect(stringCollection.GetAt(8)).To(Equal("ii"))
			Expect(stringCollection.GetAt(9)).To(Equal("jj"))
		})

		It("Should remove an element", func() {
			first := stringCollection.GetAt(0)
			firstRemoveAt := stringCollection.RemoveAt(0)

			Expect(firstRemoveAt).To(Equal(first))

			lastIndex := stringCollection.Count() - 1
			last := stringCollection.GetAt(lastIndex)
			lastRemoveAt := stringCollection.RemoveAt(lastIndex)

			Expect(lastRemoveAt).To(Equal(last))

			Expect(stringCollection.Count()).To(Equal(8))

			Expect(stringCollection.GetAt(0)).To(Equal("b"))
			Expect(stringCollection.GetAt(1)).To(Equal("c"))
			Expect(stringCollection.GetAt(2)).To(Equal("d"))
			Expect(stringCollection.GetAt(3)).To(Equal("e"))
			Expect(stringCollection.GetAt(4)).To(Equal("f"))
			Expect(stringCollection.GetAt(5)).To(Equal("g"))
			Expect(stringCollection.GetAt(6)).To(Equal("h"))
			Expect(stringCollection.GetAt(7)).To(Equal("i"))
		})

		It("Should find an element", func() {
			var index = stringCollection.FindFirst(func(_ int, element string) bool {
				return element == "e"
			})

			Expect(index).To(Equal(4))
			Expect(stringCollection.GetAt(index)).To(Equal("e"))

			firstIndex := stringCollection.FindFirst(func(_ int, element string) bool {
				return element == "b"
			})

			Expect(firstIndex).To(Equal(1))
			Expect(stringCollection.GetAt(firstIndex)).To(Equal("b"))

			lastIndex := stringCollection.FindLast(func(_ int, element string) bool {
				return element >= "a"
			})
			Expect(lastIndex).To(Equal(9))
			Expect(stringCollection.GetAt(lastIndex)).To(Equal("j"))
		})

		It("Should not find an element", func() {
			var index = stringCollection.FindFirst(func(_ int, element string) bool {
				return element == "z"
			})

			Expect(index).To(Equal(-1))

			firstIndex := stringCollection.FindFirst(func(_ int, element string) bool {
				return element == "z"
			})

			Expect(firstIndex).To(Equal(-1))

			lastIndex := stringCollection.FindLast(func(_ int, element string) bool {
				return element == "z"
			})

			Expect(lastIndex).To(Equal(-1))
		})

		It("Should try get element in range", func() {
			val, ok := stringCollection.TryGetAt(0)
			Expect(val).To(Equal("a"))
			Expect(ok).To(BeTrue())

			val, ok = stringCollection.TryGetAt(1)
			Expect(val).To(Equal("b"))
			Expect(ok).To(BeTrue())

			val, ok = stringCollection.TryGetAt(9)
			Expect(val).To(Equal("j"))
			Expect(ok).To(BeTrue())
		})

		It("Should try get element out of range", func() {
			val, ok := stringCollection.TryGetAt(-1)
			Expect(val).To(Equal(""))
			Expect(ok).To(BeFalse())

			val, ok = stringCollection.TryGetAt(10)
			Expect(val).To(Equal(""))
			Expect(ok).To(BeFalse())
		})

		It("Should try set element in range", func() {
			ok := stringCollection.TrySetAt(0, "aa")
			Expect(ok).To(BeTrue())
			Expect(stringCollection.GetAt(0)).To(Equal("aa"))

			ok = stringCollection.TrySetAt(1, "bb")
			Expect(ok).To(BeTrue())
			Expect(stringCollection.GetAt(1)).To(Equal("bb"))

			ok = stringCollection.TrySetAt(9, "jj")
			Expect(ok).To(BeTrue())
			Expect(stringCollection.GetAt(9)).To(Equal("jj"))
		})

		It("Should try set element out of range", func() {
			ok := stringCollection.TrySetAt(-1, "0")
			Expect(ok).To(BeFalse())

			ok = stringCollection.TrySetAt(10, "k")
			Expect(ok).To(BeFalse())
		})

		It("Should try remove element in range", func() {
			val, ok := stringCollection.TryRemoveAt(0)
			Expect(val).To(Equal("a"))
			Expect(ok).To(BeTrue())
			Expect(stringCollection.Count()).To(Equal(9))

			val, ok = stringCollection.TryRemoveAt(8)
			Expect(val).To(Equal("j"))
			Expect(ok).To(BeTrue())
			Expect(stringCollection.Count()).To(Equal(8))
		})

		It("Should try remove element out of range", func() {
			val, ok := stringCollection.TryRemoveAt(-1)
			Expect(val).To(Equal(""))
			Expect(ok).To(BeFalse())
			Expect(stringCollection.Count()).To(Equal(10))

			val, ok = stringCollection.TryRemoveAt(10)
			Expect(val).To(Equal(""))
			Expect(ok).To(BeFalse())
			Expect(stringCollection.Count()).To(Equal(10))
		})
	}
}

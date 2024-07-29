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

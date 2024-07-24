package linkedlist_test

import (
	"generic-collections/linkedlist"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLinkedList(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "LinkedList Suite")
}

var _ = Describe("Test LinkedList implements ICollection", func() {
	Context("Using integer", func() {
		var integerList *linkedlist.LinkedList[int]

		BeforeEach(func() {
			integerList = linkedlist.From(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

			Expect(integerList.Count()).To(Equal(10))
		})

		It("Should get an element", func() {
			Expect(integerList.GetAt(0)).To(Equal(1))
			Expect(integerList.GetAt(1)).To(Equal(2))
			Expect(integerList.GetAt(2)).To(Equal(3))
			Expect(integerList.GetAt(3)).To(Equal(4))
			Expect(integerList.GetAt(4)).To(Equal(5))
			Expect(integerList.GetAt(5)).To(Equal(6))
			Expect(integerList.GetAt(6)).To(Equal(7))
			Expect(integerList.GetAt(7)).To(Equal(8))
			Expect(integerList.GetAt(8)).To(Equal(9))
			Expect(integerList.GetAt(9)).To(Equal(10))

			Expect(integerList.NodeAt(0).Value).To(Equal(1))
			Expect(integerList.NodeAt(1).Value).To(Equal(2))
			Expect(integerList.NodeAt(2).Value).To(Equal(3))
			Expect(integerList.NodeAt(3).Value).To(Equal(4))
			Expect(integerList.NodeAt(4).Value).To(Equal(5))
			Expect(integerList.NodeAt(5).Value).To(Equal(6))
			Expect(integerList.NodeAt(6).Value).To(Equal(7))
			Expect(integerList.NodeAt(7).Value).To(Equal(8))
			Expect(integerList.NodeAt(8).Value).To(Equal(9))
			Expect(integerList.NodeAt(9).Value).To(Equal(10))
		})

		It("Should set an element", func() {
			integerList.SetAt(0, 11)
			integerList.SetAt(1, 12)
			integerList.SetAt(2, 13)
			integerList.SetAt(3, 14)
			integerList.SetAt(4, 15)
			integerList.SetAt(5, 16)
			integerList.SetAt(6, 17)
			integerList.SetAt(7, 18)
			integerList.SetAt(8, 19)
			integerList.SetAt(9, 20)

			Expect(integerList.GetAt(0)).To(Equal(11))
			Expect(integerList.GetAt(1)).To(Equal(12))
			Expect(integerList.GetAt(2)).To(Equal(13))
			Expect(integerList.GetAt(3)).To(Equal(14))
			Expect(integerList.GetAt(4)).To(Equal(15))
			Expect(integerList.GetAt(5)).To(Equal(16))
			Expect(integerList.GetAt(6)).To(Equal(17))
			Expect(integerList.GetAt(7)).To(Equal(18))
			Expect(integerList.GetAt(8)).To(Equal(19))
			Expect(integerList.GetAt(9)).To(Equal(20))
		})

		It("Should remove an element", func() {
			first := integerList.GetAt(0)
			firstRemoveAt := integerList.RemoveAt(0)

			Expect(firstRemoveAt).To(Equal(first))

			lastIndex := integerList.Count() - 1
			last := integerList.GetAt(lastIndex)
			lastRemoveAt := integerList.RemoveAt(lastIndex)

			Expect(lastRemoveAt).To(Equal(last))

			Expect(integerList.Count()).To(Equal(8))

			Expect(integerList.GetAt(0)).To(Equal(2))
			Expect(integerList.GetAt(1)).To(Equal(3))
			Expect(integerList.GetAt(2)).To(Equal(4))
			Expect(integerList.GetAt(3)).To(Equal(5))
			Expect(integerList.GetAt(4)).To(Equal(6))
			Expect(integerList.GetAt(5)).To(Equal(7))
			Expect(integerList.GetAt(6)).To(Equal(8))
			Expect(integerList.GetAt(7)).To(Equal(9))

			Expect(integerList.Head.Value).To(Equal(2))
			Expect(integerList.Tail.Value).To(Equal(9))
		})

		It("Should find an element", func() {
			var index = integerList.Find(func(element int) bool {
				return element >= 5
			})

			Expect(index).To(Equal(4))
			Expect(integerList.GetAt(index)).To(Equal(5))
		})

		It("Should not find an element", func() {
			var index = integerList.Find(func(element int) bool {
				return element == 11
			})

			Expect(index).To(Equal(-1))
		})

		It("Should try get element in range", func() {
			val, ok := integerList.TryGetAt(0)
			Expect(val).To(Equal(1))
			Expect(ok).To(BeTrue())

			val, ok = integerList.TryGetAt(1)
			Expect(val).To(Equal(2))
			Expect(ok).To(BeTrue())

			val, ok = integerList.TryGetAt(9)
			Expect(val).To(Equal(10))
			Expect(ok).To(BeTrue())
		})

		It("Should try get element out of range", func() {
			val, ok := integerList.TryGetAt(-1)
			Expect(val).To(Equal(0))
			Expect(ok).To(BeFalse())

			val, ok = integerList.TryGetAt(10)
			Expect(val).To(Equal(0))
			Expect(ok).To(BeFalse())
		})

		It("Should try set element in range", func() {
			ok := integerList.TrySetAt(0, 11)
			Expect(ok).To(BeTrue())
			Expect(integerList.GetAt(0)).To(Equal(11))

			ok = integerList.TrySetAt(1, 12)
			Expect(ok).To(BeTrue())
			Expect(integerList.GetAt(1)).To(Equal(12))

			ok = integerList.TrySetAt(9, 20)
			Expect(ok).To(BeTrue())
			Expect(integerList.GetAt(9)).To(Equal(20))
		})

		It("Should try set element out of range", func() {
			ok := integerList.TrySetAt(-1, 0)
			Expect(ok).To(BeFalse())

			ok = integerList.TrySetAt(10, 21)
			Expect(ok).To(BeFalse())
		})

		It("Should try remove element in range", func() {
			val, ok := integerList.TryRemoveAt(0)
			Expect(val).To(Equal(1))
			Expect(ok).To(BeTrue())
			Expect(integerList.Count()).To(Equal(9))

			val, ok = integerList.TryRemoveAt(8)
			Expect(val).To(Equal(10))
			Expect(ok).To(BeTrue())
			Expect(integerList.Count()).To(Equal(8))
		})

		It("Should try remove element out of range", func() {
			val, ok := integerList.TryRemoveAt(-1)
			Expect(val).To(Equal(0))
			Expect(ok).To(BeFalse())
			Expect(integerList.Count()).To(Equal(10))

			val, ok = integerList.TryRemoveAt(10)
			Expect(val).To(Equal(0))
			Expect(ok).To(BeFalse())
			Expect(integerList.Count()).To(Equal(10))
		})
	})

	Context("Using string", func() {
		var stringList *linkedlist.LinkedList[string]

		BeforeEach(func() {
			stringList = linkedlist.From("a", "b", "c", "d", "e", "f", "g", "h", "i", "j")

			Expect(stringList.Count()).To(Equal(10))
		})

		It("Should get an element", func() {
			Expect(stringList.GetAt(0)).To(Equal("a"))
			Expect(stringList.GetAt(1)).To(Equal("b"))
			Expect(stringList.GetAt(2)).To(Equal("c"))
			Expect(stringList.GetAt(3)).To(Equal("d"))
			Expect(stringList.GetAt(4)).To(Equal("e"))
			Expect(stringList.GetAt(5)).To(Equal("f"))
			Expect(stringList.GetAt(6)).To(Equal("g"))
			Expect(stringList.GetAt(7)).To(Equal("h"))
			Expect(stringList.GetAt(8)).To(Equal("i"))
			Expect(stringList.GetAt(9)).To(Equal("j"))
		})

		It("Should set an element", func() {
			stringList.SetAt(0, "aa")
			stringList.SetAt(1, "bb")
			stringList.SetAt(2, "cc")
			stringList.SetAt(3, "dd")
			stringList.SetAt(4, "ee")
			stringList.SetAt(5, "ff")
			stringList.SetAt(6, "gg")
			stringList.SetAt(7, "hh")
			stringList.SetAt(8, "ii")
			stringList.SetAt(9, "jj")

			Expect(stringList.GetAt(0)).To(Equal("aa"))
			Expect(stringList.GetAt(1)).To(Equal("bb"))
			Expect(stringList.GetAt(2)).To(Equal("cc"))
			Expect(stringList.GetAt(3)).To(Equal("dd"))
			Expect(stringList.GetAt(4)).To(Equal("ee"))
			Expect(stringList.GetAt(5)).To(Equal("ff"))
			Expect(stringList.GetAt(6)).To(Equal("gg"))
			Expect(stringList.GetAt(7)).To(Equal("hh"))
			Expect(stringList.GetAt(8)).To(Equal("ii"))
			Expect(stringList.GetAt(9)).To(Equal("jj"))
		})

		It("Should remove an element", func() {
			first := stringList.GetAt(0)
			firstRemoveAt := stringList.RemoveAt(0)

			Expect(firstRemoveAt).To(Equal(first))

			lastIndex := stringList.Count() - 1
			last := stringList.GetAt(lastIndex)
			lastRemoveAt := stringList.RemoveAt(lastIndex)

			Expect(lastRemoveAt).To(Equal(last))

			Expect(stringList.Count()).To(Equal(8))

			Expect(stringList.GetAt(0)).To(Equal("b"))
			Expect(stringList.GetAt(1)).To(Equal("c"))
			Expect(stringList.GetAt(2)).To(Equal("d"))
			Expect(stringList.GetAt(3)).To(Equal("e"))
			Expect(stringList.GetAt(4)).To(Equal("f"))
			Expect(stringList.GetAt(5)).To(Equal("g"))
			Expect(stringList.GetAt(6)).To(Equal("h"))
			Expect(stringList.GetAt(7)).To(Equal("i"))

			Expect(stringList.Head.Value).To(Equal("b"))
			Expect(stringList.Tail.Value).To(Equal("i"))
		})

		It("Should find an element", func() {
			var index = stringList.Find(func(element string) bool {
				return element == "e"
			})

			Expect(index).To(Equal(4))
			Expect(stringList.GetAt(index)).To(Equal("e"))
		})

		It("Should not find an element", func() {
			var index = stringList.Find(func(element string) bool {
				return element == "z"
			})

			Expect(index).To(Equal(-1))
		})

		It("Should try get element in range", func() {
			val, ok := stringList.TryGetAt(0)
			Expect(val).To(Equal("a"))
			Expect(ok).To(BeTrue())

			val, ok = stringList.TryGetAt(1)
			Expect(val).To(Equal("b"))
			Expect(ok).To(BeTrue())

			val, ok = stringList.TryGetAt(9)
			Expect(val).To(Equal("j"))
			Expect(ok).To(BeTrue())
		})

		It("Should try get element out of range", func() {
			val, ok := stringList.TryGetAt(-1)
			Expect(val).To(Equal(""))
			Expect(ok).To(BeFalse())

			val, ok = stringList.TryGetAt(10)
			Expect(val).To(Equal(""))
			Expect(ok).To(BeFalse())
		})

		It("Should try set element in range", func() {
			ok := stringList.TrySetAt(0, "aa")
			Expect(ok).To(BeTrue())
			Expect(stringList.GetAt(0)).To(Equal("aa"))

			ok = stringList.TrySetAt(1, "bb")
			Expect(ok).To(BeTrue())
			Expect(stringList.GetAt(1)).To(Equal("bb"))

			ok = stringList.TrySetAt(9, "jj")
			Expect(ok).To(BeTrue())
			Expect(stringList.GetAt(9)).To(Equal("jj"))
		})

		It("Should try set element out of range", func() {
			ok := stringList.TrySetAt(-1, "0")
			Expect(ok).To(BeFalse())

			ok = stringList.TrySetAt(10, "k")
			Expect(ok).To(BeFalse())
		})

		It("Should try remove element in range", func() {
			val, ok := stringList.TryRemoveAt(0)
			Expect(val).To(Equal("a"))
			Expect(ok).To(BeTrue())
			Expect(stringList.Count()).To(Equal(9))

			val, ok = stringList.TryRemoveAt(8)
			Expect(val).To(Equal("j"))
			Expect(ok).To(BeTrue())
			Expect(stringList.Count()).To(Equal(8))
		})

		It("Should try remove element out of range", func() {
			val, ok := stringList.TryRemoveAt(-1)
			Expect(val).To(Equal(""))
			Expect(ok).To(BeFalse())
			Expect(stringList.Count()).To(Equal(10))

			val, ok = stringList.TryRemoveAt(10)
			Expect(val).To(Equal(""))
			Expect(ok).To(BeFalse())
			Expect(stringList.Count()).To(Equal(10))
		})
	})

	Context("Try panic", func() {
		var integerList *linkedlist.LinkedList[int]

		BeforeEach(func() {
			integerList = linkedlist.From(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

		})

		It("Should panic", func() {
			Expect(func() {
				integerList.GetAt(-1)
			}).To(Panic())

			Expect(func() {
				integerList.NodeAt(-1)
			}).To(Panic())

			Expect(func() {
				integerList.GetAt(10)
			}).To(Panic())

			Expect(func() {
				integerList.SetAt(-1, 0)
			}).To(Panic())

			Expect(func() {
				integerList.SetAt(10, 0)
			}).To(Panic())

			Expect(func() {
				integerList.RemoveAt(-1)
			}).To(Panic())

			Expect(func() {
				integerList.RemoveAt(10)
			}).To(Panic())
		})
	})
})

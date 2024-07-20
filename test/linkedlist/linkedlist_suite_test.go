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
			Expect(integerList.Get(0)).To(Equal(1))
			Expect(integerList.Get(1)).To(Equal(2))
			Expect(integerList.Get(2)).To(Equal(3))
			Expect(integerList.Get(3)).To(Equal(4))
			Expect(integerList.Get(4)).To(Equal(5))
			Expect(integerList.Get(5)).To(Equal(6))
			Expect(integerList.Get(6)).To(Equal(7))
			Expect(integerList.Get(7)).To(Equal(8))
			Expect(integerList.Get(8)).To(Equal(9))
			Expect(integerList.Get(9)).To(Equal(10))
		})

		It("Should set an element", func() {
			integerList.Set(0, 11)
			integerList.Set(1, 12)
			integerList.Set(2, 13)
			integerList.Set(3, 14)
			integerList.Set(4, 15)
			integerList.Set(5, 16)
			integerList.Set(6, 17)
			integerList.Set(7, 18)
			integerList.Set(8, 19)
			integerList.Set(9, 20)

			Expect(integerList.Get(0)).To(Equal(11))
			Expect(integerList.Get(1)).To(Equal(12))
			Expect(integerList.Get(2)).To(Equal(13))
			Expect(integerList.Get(3)).To(Equal(14))
			Expect(integerList.Get(4)).To(Equal(15))
			Expect(integerList.Get(5)).To(Equal(16))
			Expect(integerList.Get(6)).To(Equal(17))
			Expect(integerList.Get(7)).To(Equal(18))
			Expect(integerList.Get(8)).To(Equal(19))
			Expect(integerList.Get(9)).To(Equal(20))
		})

		It("Should remove an element", func() {
			first := integerList.Get(0)
			firstRemove := integerList.Remove(0)

			Expect(firstRemove).To(Equal(first))

			lastIndex := integerList.Count() - 1
			last := integerList.Get(lastIndex)
			lastRemove := integerList.Remove(lastIndex)

			Expect(lastRemove).To(Equal(last))

			Expect(integerList.Count()).To(Equal(8))

			Expect(integerList.Get(0)).To(Equal(2))
			Expect(integerList.Get(1)).To(Equal(3))
			Expect(integerList.Get(2)).To(Equal(4))
			Expect(integerList.Get(3)).To(Equal(5))
			Expect(integerList.Get(4)).To(Equal(6))
			Expect(integerList.Get(5)).To(Equal(7))
			Expect(integerList.Get(6)).To(Equal(8))
			Expect(integerList.Get(7)).To(Equal(9))

			Expect(integerList.Head.Value).To(Equal(2))
			Expect(integerList.Tail.Value).To(Equal(9))
		})

		It("Should find an element", func() {
			var index = integerList.Find(func(element int) bool {
				return element >= 5
			})

			Expect(index).To(Equal(4))
			Expect(integerList.Get(index)).To(Equal(5))
		})

		It("Should not find an element", func() {
			var index = integerList.Find(func(element int) bool {
				return element == 11
			})

			Expect(index).To(Equal(-1))
		})
	})

	Context("Using string", func() {
		var stringList *linkedlist.LinkedList[string]

		BeforeEach(func() {
			stringList = linkedlist.From("a", "b", "c", "d", "e", "f", "g", "h", "i", "j")

			Expect(stringList.Count()).To(Equal(10))
		})

		It("Should get an element", func() {
			Expect(stringList.Get(0)).To(Equal("a"))
			Expect(stringList.Get(1)).To(Equal("b"))
			Expect(stringList.Get(2)).To(Equal("c"))
			Expect(stringList.Get(3)).To(Equal("d"))
			Expect(stringList.Get(4)).To(Equal("e"))
			Expect(stringList.Get(5)).To(Equal("f"))
			Expect(stringList.Get(6)).To(Equal("g"))
			Expect(stringList.Get(7)).To(Equal("h"))
			Expect(stringList.Get(8)).To(Equal("i"))
			Expect(stringList.Get(9)).To(Equal("j"))
		})

		It("Should set an element", func() {
			stringList.Set(0, "aa")
			stringList.Set(1, "bb")
			stringList.Set(2, "cc")
			stringList.Set(3, "dd")
			stringList.Set(4, "ee")
			stringList.Set(5, "ff")
			stringList.Set(6, "gg")
			stringList.Set(7, "hh")
			stringList.Set(8, "ii")
			stringList.Set(9, "jj")

			Expect(stringList.Get(0)).To(Equal("aa"))
			Expect(stringList.Get(1)).To(Equal("bb"))
			Expect(stringList.Get(2)).To(Equal("cc"))
			Expect(stringList.Get(3)).To(Equal("dd"))
			Expect(stringList.Get(4)).To(Equal("ee"))
			Expect(stringList.Get(5)).To(Equal("ff"))
			Expect(stringList.Get(6)).To(Equal("gg"))
			Expect(stringList.Get(7)).To(Equal("hh"))
			Expect(stringList.Get(8)).To(Equal("ii"))
			Expect(stringList.Get(9)).To(Equal("jj"))
		})

		It("Should remove an element", func() {
			first := stringList.Get(0)
			firstRemove := stringList.Remove(0)

			Expect(firstRemove).To(Equal(first))

			lastIndex := stringList.Count() - 1
			last := stringList.Get(lastIndex)
			lastRemove := stringList.Remove(lastIndex)

			Expect(lastRemove).To(Equal(last))

			Expect(stringList.Count()).To(Equal(8))

			Expect(stringList.Get(0)).To(Equal("b"))
			Expect(stringList.Get(1)).To(Equal("c"))
			Expect(stringList.Get(2)).To(Equal("d"))
			Expect(stringList.Get(3)).To(Equal("e"))
			Expect(stringList.Get(4)).To(Equal("f"))
			Expect(stringList.Get(5)).To(Equal("g"))
			Expect(stringList.Get(6)).To(Equal("h"))
			Expect(stringList.Get(7)).To(Equal("i"))

			Expect(stringList.Head.Value).To(Equal("b"))
			Expect(stringList.Tail.Value).To(Equal("i"))
		})

		It("Should find an element", func() {
			var index = stringList.Find(func(element string) bool {
				return element == "e"
			})

			Expect(index).To(Equal(4))
			Expect(stringList.Get(index)).To(Equal("e"))
		})

		It("Should not find an element", func() {
			var index = stringList.Find(func(element string) bool {
				return element == "z"
			})

			Expect(index).To(Equal(-1))
		})
	})

	Context("Try panic", func() {
		var integerList *linkedlist.LinkedList[int]

		BeforeEach(func() {
			integerList = linkedlist.From(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

		})

		It("Should panic", func() {
			Expect(func() {
				integerList.Get(-1)
			}).To(Panic())

			Expect(func() {
				integerList.Get(10)
			}).To(Panic())

			Expect(func() {
				integerList.Set(-1, 0)
			}).To(Panic())

			Expect(func() {
				integerList.Set(10, 0)
			}).To(Panic())

			Expect(func() {
				integerList.Remove(-1)
			}).To(Panic())

			Expect(func() {
				integerList.Remove(10)
			}).To(Panic())
		})
	})
})

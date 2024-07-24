package list

import (
	"generic-collections/list"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"strings"
)

var _ = Describe("Test List implements ICollection", func() {

	Context("Using integer", func() {
		var integerList *list.List[int]

		BeforeEach(func() {
			integerList = list.From(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

			Expect(integerList.Count()).To(Equal(10))
			Expect(list.IsList[int](integerList)).To(BeTrue())
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
			integerList.RemoveAt(0)
			integerList.RemoveAt(0)
			integerList.RemoveAt(0)
			integerList.RemoveAt(0)
			integerList.RemoveAt(0)
			integerList.RemoveAt(0)
			integerList.RemoveAt(0)
			integerList.RemoveAt(0)
			integerList.RemoveAt(0)
			integerList.RemoveAt(0)

			Expect(integerList.Count()).To(Equal(0))

			Expect(func() {
				integerList.RemoveAt(-1)
			}).To(Panic())

			Expect(func() {
				integerList.RemoveAt(0)
			}).To(Panic())
		})

		It("Should convert to slice", func() {
			Expect(integerList.ToSlice()).To(Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
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
			var item, ok = integerList.TryGetAt(5)

			Expect(ok).To(BeTrue())
			Expect(item).To(Equal(6))
		})

		It("Should try to get element out of range", func() {
			var item, ok = integerList.TryGetAt(10)

			Expect(ok).To(BeFalse())
			Expect(item).To(Equal(0))
		})

		It("Should try set element in range", func() {
			var ok = integerList.TrySetAt(5, 100)

			Expect(ok).To(BeTrue())
			Expect(integerList.GetAt(5)).To(Equal(100))
		})

		It("Should try to set element out of range", func() {
			var ok = integerList.TrySetAt(10, 100)

			Expect(ok).To(BeFalse())
			Expect(integerList.Count()).To(Equal(10))
		})

		It("Should try remove element in range", func() {
			var item, ok = integerList.TryRemoveAt(5)

			Expect(ok).To(BeTrue())
			Expect(item).To(Equal(6))
			Expect(integerList.Count()).To(Equal(9))
		})

		It("Should try to remove element out of range", func() {
			var item, ok = integerList.TryRemoveAt(10)

			Expect(ok).To(BeFalse())
			Expect(item).To(Equal(0))
			Expect(integerList.Count()).To(Equal(10))
		})
	})

	Context("Using string", func() {
		var stringList *list.List[string]

		BeforeEach(func() {
			stringList = list.From(
				"Apple",
				"Banana",
				"Cherry",
				"Dates",
				"Elderberry",
				"Fig",
				"Grape",
				"Honeydew",
				"Jackfruit",
				"Kiwi",
			)

			Expect(stringList.Count()).To(Equal(10))
		})

		It("Should get elements", func() {
			Expect(stringList.GetAt(0)).To(Equal("Apple"))
			Expect(stringList.GetAt(1)).To(Equal("Banana"))
			Expect(stringList.GetAt(2)).To(Equal("Cherry"))
			Expect(stringList.GetAt(3)).To(Equal("Dates"))
			Expect(stringList.GetAt(4)).To(Equal("Elderberry"))
			Expect(stringList.GetAt(5)).To(Equal("Fig"))
			Expect(stringList.GetAt(6)).To(Equal("Grape"))
			Expect(stringList.GetAt(7)).To(Equal("Honeydew"))
			Expect(stringList.GetAt(8)).To(Equal("Jackfruit"))
			Expect(stringList.GetAt(9)).To(Equal("Kiwi"))
		})

		It("Should set elements", func() {
			stringList.SetAt(0, "Lemon")
			stringList.SetAt(1, "Mango")
			stringList.SetAt(2, "Nectarine")
			stringList.SetAt(3, "Orange")
			stringList.SetAt(4, "Papaya")
			stringList.SetAt(5, "Quince")

			Expect(stringList.GetAt(0)).To(Equal("Lemon"))
			Expect(stringList.GetAt(1)).To(Equal("Mango"))
			Expect(stringList.GetAt(2)).To(Equal("Nectarine"))
			Expect(stringList.GetAt(3)).To(Equal("Orange"))
			Expect(stringList.GetAt(4)).To(Equal("Papaya"))
			Expect(stringList.GetAt(5)).To(Equal("Quince"))
		})

		It("Should convert to slice", func() {
			Expect(stringList.ToSlice()).To(Equal([]string{"Apple", "Banana", "Cherry", "Dates", "Elderberry", "Fig", "Grape", "Honeydew", "Jackfruit", "Kiwi"}))
		})

		It("Should find an element", func() {
			var index = stringList.Find(func(element string) bool {
				return element == "Grape"
			})

			Expect(index).To(Equal(6))
			Expect(stringList.GetAt(index)).To(Equal("Grape"))

			index = stringList.Find(func(element string) bool {
				return strings.Contains(element, "r")
			})

			Expect(index).To(Equal(2))
			Expect(stringList.GetAt(index)).To(Equal("Cherry"))
		})

		It("Should not find an element", func() {
			var index = stringList.Find(func(element string) bool {
				return element == "Lime"
			})

			Expect(index).To(Equal(-1))

			index = stringList.Find(func(element string) bool {
				return strings.Contains(element, "z")
			})

			Expect(index).To(Equal(-1))
		})

		It("Should try get element in range", func() {
			var item, ok = stringList.TryGetAt(5)

			Expect(ok).To(BeTrue())
			Expect(item).To(Equal("Fig"))
		})

		It("Should try to get element out of range", func() {
			var item, ok = stringList.TryGetAt(10)

			Expect(ok).To(BeFalse())
			Expect(item).To(Equal(""))
		})

		It("Should try set element in range", func() {
			var ok = stringList.TrySetAt(5, "Lemon")

			Expect(ok).To(BeTrue())
			Expect(stringList.GetAt(5)).To(Equal("Lemon"))
		})

		It("Should try to set element out of range", func() {
			var ok = stringList.TrySetAt(10, "Lemon")

			Expect(ok).To(BeFalse())
		})

		It("Should try remove element in range", func() {
			var item, ok = stringList.TryRemoveAt(5)

			Expect(ok).To(BeTrue())
			Expect(item).To(Equal("Fig"))
			Expect(stringList.Count()).To(Equal(9))
		})

		It("Should try to remove element out of range", func() {
			var item, ok = stringList.TryRemoveAt(10)

			Expect(ok).To(BeFalse())
			Expect(item).To(Equal(""))
			Expect(stringList.Count()).To(Equal(10))
		})
	})

	Context("Try panic", func() {
		var floatList *list.List[float64]

		BeforeEach(func() {
			floatList = list.From(1.1, 2.2, 3.3, 4.4, 5.5)
		})

		It("Should panic when index out of range", func() {
			Expect(func() {
				floatList.GetAt(5)
				floatList.SetAt(5, 6.6)
			}).To(Panic())
		})

		It("Should panic when index is negative", func() {
			Expect(func() {
				floatList.GetAt(-1)
				floatList.SetAt(-1, 6.6)
			}).To(Panic())
		})
	})
})

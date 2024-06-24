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

		It("Should convert to slice", func() {
			Expect(integerList.ToSlice()).To(Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
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
			Expect(stringList.Get(0)).To(Equal("Apple"))
			Expect(stringList.Get(1)).To(Equal("Banana"))
			Expect(stringList.Get(2)).To(Equal("Cherry"))
			Expect(stringList.Get(3)).To(Equal("Dates"))
			Expect(stringList.Get(4)).To(Equal("Elderberry"))
			Expect(stringList.Get(5)).To(Equal("Fig"))
			Expect(stringList.Get(6)).To(Equal("Grape"))
			Expect(stringList.Get(7)).To(Equal("Honeydew"))
			Expect(stringList.Get(8)).To(Equal("Jackfruit"))
			Expect(stringList.Get(9)).To(Equal("Kiwi"))
		})

		It("Should set elements", func() {
			stringList.Set(0, "Lemon")
			stringList.Set(1, "Mango")
			stringList.Set(2, "Nectarine")
			stringList.Set(3, "Orange")
			stringList.Set(4, "Papaya")
			stringList.Set(5, "Quince")

			Expect(stringList.Get(0)).To(Equal("Lemon"))
			Expect(stringList.Get(1)).To(Equal("Mango"))
			Expect(stringList.Get(2)).To(Equal("Nectarine"))
			Expect(stringList.Get(3)).To(Equal("Orange"))
			Expect(stringList.Get(4)).To(Equal("Papaya"))
			Expect(stringList.Get(5)).To(Equal("Quince"))
		})

		It("Should convert to slice", func() {
			Expect(stringList.ToSlice()).To(Equal([]string{"Apple", "Banana", "Cherry", "Dates", "Elderberry", "Fig", "Grape", "Honeydew", "Jackfruit", "Kiwi"}))
		})

		It("Should find an element", func() {
			var index = stringList.Find(func(element string) bool {
				return element == "Grape"
			})

			Expect(index).To(Equal(6))
			Expect(stringList.Get(index)).To(Equal("Grape"))

			index = stringList.Find(func(element string) bool {
				return strings.Contains(element, "r")
			})

			Expect(index).To(Equal(2))
			Expect(stringList.Get(index)).To(Equal("Cherry"))
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
	})

	Context("Try panic", func() {
		var floatList *list.List[float64]

		BeforeEach(func() {
			floatList = list.From(1.1, 2.2, 3.3, 4.4, 5.5)
		})

		It("Should panic when index out of range", func() {
			Expect(func() {
				floatList.Get(5)
				floatList.Set(5, 6.6)
			}).To(Panic())
		})

		It("Should panic when index is negative", func() {
			Expect(func() {
				floatList.Get(-1)
				floatList.Set(-1, 6.6)
			}).To(Panic())
		})
	})
})

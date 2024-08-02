package list

import (
	"generic-collections/list"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test List implements ICollection", func() {

	Context("Using integer", func() {
		var integerList *list.List[int]

		BeforeEach(func() {
			integerList = list.From(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

			Expect(integerList.Count()).To(Equal(10))
			Expect(list.IsList[int](integerList)).To(BeTrue())
		})

		It("Should convert to slice", func() {
			Expect(integerList.ToSlice()).To(Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
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

		It("Should convert to slice", func() {
			Expect(stringList.ToSlice()).To(Equal([]string{"Apple", "Banana", "Cherry", "Dates", "Elderberry", "Fig", "Grape", "Honeydew", "Jackfruit", "Kiwi"}))
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

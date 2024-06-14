package list_test

import (
	"generic-collections/list"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestList(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "List Suite")
}

var _ = Describe("Test List implements ICollection", func() {

	When("Using integer list", func() {
		var integerList *list.List[int]

		BeforeEach(func() {
			integerList = list.From(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

			Expect(integerList.Count()).To(Equal(10))
		})

		It("Should add an element", func() {
			integerList.Add(11)

			Expect(integerList.Count()).To(Equal(11))
		})

		It("Should add all elements", func() {
			integerList.AddAll(list.From(11, 12, 13, 14, 15))

			Expect(integerList.Count()).To(Equal(15))
		})

		It("Should clear all elements", func() {
			integerList.Clear()

			Expect(integerList.Count()).To(Equal(0))
			Expect(integerList.IsEmpty(), BeTrue())
		})

		It("Should remove an element", func() {
			integerList.Remove(5)

			Expect(integerList.Count()).To(Equal(9))
			Expect(integerList.IsEmpty(), BeFalse())
		})

		It("Should remove all elements", func() {
			integerList.RemoveAll(list.From(1, 2, 3, 4, 5))

			Expect(integerList.Count()).To(Equal(5))
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

		It("Should check if contains an element", func() {
			Expect(integerList.Contains(5)).To(BeTrue())
			Expect(integerList.Contains(15)).To(BeFalse())
		})

		It("Should check if contains all elements", func() {
			Expect(integerList.ContainsAll(list.From(1, 2, 3, 4, 5))).To(BeTrue())
			Expect(integerList.ContainsAll(list.From(1, 2, 3, 4, 15))).To(BeFalse())
		})

		It("Should filter elements", func() {
			filtered := integerList.Filter(func(element int) bool {
				return element > 5
			})

			Expect(filtered.ToSlice()).To(Equal([]int{6, 7, 8, 9, 10}))
			Expect(filtered.Count()).To(Equal(5))
		})

		It("Should iterate over elements", func() {
			var sum int
			integerList.ForEach(func(_ int, element int) {
				sum += element
			})

			Expect(sum).To(Equal(55))
		})
	})

})

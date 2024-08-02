package linkedlist_test

import (
	"generic-collections/linkedlist"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type Company struct {
	Name    string
	Country string
	Value   float64
}

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

		It("Should clone a node", func() {
			node := integerList.NodeAt(0)
			cloned := node.Clone()

			Expect(cloned.Value).To(Equal(node.Value))
			Expect(cloned.Next).To(BeNil())
			Expect(cloned).ToNot(Equal(node))
		})
	})

	Context("Using string", func() {
		var stringList *linkedlist.LinkedList[string]

		BeforeEach(func() {
			stringList = linkedlist.From("a", "b", "c", "d", "e", "f", "g", "h", "i", "j")

			Expect(stringList.Count()).To(Equal(10))
		})

	})

	Context("Using struct", func() {
		var companyList *linkedlist.LinkedList[Company]

		BeforeEach(func() {
			companyList = linkedlist.From(
				Company{"Company A", "Country A", 1000},
				Company{"Company B", "Country B", 2000},
				Company{"Company C", "Country C", 3000},
				Company{"Company D", "Country D", 4000},
				Company{"Company E", "Country E", 5000},
				Company{"Company F", "Country F", 6000},
				Company{"Company G", "Country G", 7000},
				Company{"Company H", "Country H", 8000},
				Company{"Company I", "Country I", 9000},
				Company{"Company J", "Country J", 10000},
			)

			Expect(companyList.Count()).To(Equal(10))
		})

		It("Should map", func() {
			mapped := companyList.Map(func(index int, company Company) any {
				return company.Name
			})

			Expect(mapped.Count()).To(Equal(10))
			Expect(mapped.GetAt(0)).To(Equal("Company A"))
			Expect(mapped.GetAt(1)).To(Equal("Company B"))
			Expect(mapped.GetAt(2)).To(Equal("Company C"))
			Expect(mapped.GetAt(3)).To(Equal("Company D"))
			Expect(mapped.GetAt(4)).To(Equal("Company E"))
			Expect(mapped.GetAt(5)).To(Equal("Company F"))
			Expect(mapped.GetAt(6)).To(Equal("Company G"))
			Expect(mapped.GetAt(7)).To(Equal("Company H"))
			Expect(mapped.GetAt(8)).To(Equal("Company I"))
			Expect(mapped.GetAt(9)).To(Equal("Company J"))
		})

		It("Should reduce", func() {
			totalValue := companyList.Reduce(func(accumulator any, company Company) any {
				return accumulator.(float64) + company.Value
			}, 0.0).(float64)

			Expect(totalValue).To(Equal(55000.0))
		})

		It("Should group by", func() {
			groups := companyList.GroupBy(func(company Company) any {
				if company.Value < 5000 {
					return "Small"
				} else if company.Value < 8000 {
					return "Medium"
				} else {
					return "Large"
				}
			})

			Expect(groups.Count()).To(Equal(3))

			smallCompanies := groups.Get("Small")
			Expect(smallCompanies.Count()).To(Equal(4))

			mediumCompanies := groups.Get("Medium")
			Expect(mediumCompanies.Count()).To(Equal(3))

			largeCompanies := groups.Get("Large")
			Expect(largeCompanies.Count()).To(Equal(3))
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

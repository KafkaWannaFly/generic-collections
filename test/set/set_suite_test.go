package set_test

import (
	"generic-collections/set"
	"strconv"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSet(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Set Suite")
}

type Person struct {
	Name    string
	Age     int
	friends *set.Set[Person]
}

func (receiver Person) GetHashCode() string {
	return receiver.Name + " " + strconv.Itoa(receiver.Age)
}

var _ = Describe("Set Specific Test", func() {
	When("Using built-in type", func() {
		var set1 *set.Set[int]
		var set2 *set.Set[int]

		BeforeEach(func() {
			set1 = set.From(1, 2, 3, 4, 5)
			set2 = set.From(2, 3, 4, 5, 6)

			Expect(set1.Count()).To(Equal(5))
			Expect(set2.Count()).To(Equal(5))
		})

		It("Should union", func() {
			var union = set1.Union(set2)
			Expect(union.Count()).To(Equal(6))

			Expect(union.Has(1)).To(BeTrue())
			Expect(union.Has(2)).To(BeTrue())
			Expect(union.Has(3)).To(BeTrue())
			Expect(union.Has(4)).To(BeTrue())
			Expect(union.Has(5)).To(BeTrue())
			Expect(union.Has(6)).To(BeTrue())

			var union2 = set2.Union(set1)
			Expect(union2.Count()).To(Equal(6))

			Expect(union2.ToSlice()).To(ConsistOf(1, 2, 3, 4, 5, 6))
		})

		It("Should intersect", func() {
			var intersection = set1.Intersect(set2)
			Expect(intersection.Count()).To(Equal(4))

			Expect(intersection.Has(2)).To(BeTrue())
			Expect(intersection.Has(3)).To(BeTrue())
			Expect(intersection.Has(4)).To(BeTrue())
			Expect(intersection.Has(5)).To(BeTrue())

			Expect(intersection.Has(1)).To(BeFalse())
			Expect(intersection.Has(6)).To(BeFalse())

			var intersection2 = set2.Intersect(set1)
			Expect(intersection2.Count()).To(Equal(4))

			Expect(intersection2.ToSlice()).To(ConsistOf(2, 3, 4, 5))
		})

		It("Should difference", func() {
			var difference = set1.Difference(set2)
			Expect(difference.Count()).To(Equal(1))

			Expect(difference.Has(1)).To(BeTrue())

			Expect(difference.Has(2)).To(BeFalse())
			Expect(difference.Has(3)).To(BeFalse())
			Expect(difference.Has(4)).To(BeFalse())
			Expect(difference.Has(5)).To(BeFalse())
			Expect(difference.Has(6)).To(BeFalse())

			var difference2 = set2.Difference(set1)
			Expect(difference2.Count()).To(Equal(1))

			Expect(difference2.Has(6)).To(BeTrue())
		})

		It("Should symmetric difference", func() {
			var symmetricDifference = set1.SymmetricDifference(set2)
			Expect(symmetricDifference.Count()).To(Equal(2))

			Expect(symmetricDifference.Has(1)).To(BeTrue())
			Expect(symmetricDifference.Has(6)).To(BeTrue())

			Expect(symmetricDifference.Has(2)).To(BeFalse())
			Expect(symmetricDifference.Has(3)).To(BeFalse())
			Expect(symmetricDifference.Has(4)).To(BeFalse())
			Expect(symmetricDifference.Has(5)).To(BeFalse())

			var symmetricDifference2 = set2.SymmetricDifference(set1)
			Expect(symmetricDifference2.Count()).To(Equal(2))

			Expect(symmetricDifference2.ToSlice()).To(ConsistOf(1, 6))
		})
	})

	When("Using custom type", func() {
		var set1 *set.Set[Person]
		var set2 *set.Set[Person]

		var p1 = Person{Name: "Alice", Age: 20}
		var p2 = Person{Name: "Bob", Age: 25}
		var p3 = Person{Name: "Charlie", Age: 30}
		var p4 = Person{Name: "David", Age: 35}
		var p5 = Person{Name: "Eve", Age: 40}

		BeforeEach(func() {
			p1.friends = set.From(p2, p3, p4, p5)
			p2.friends = set.From(p1, p3, p4, p5)
			p3.friends = set.From(p1, p2, p4, p5)
			p4.friends = set.From(p1, p2, p3, p5)
			p5.friends = set.From(p1, p2, p3, p4)

			set1 = set.From(p1, p2, p3, p4)
			set2 = set.From(p2, p3, p4, p5)
		})

		It("Should check data type", func() {
			Expect(set.IsSet[Person](set1)).To(BeTrue())
			Expect(set.IsSet[Person](set2)).To(BeTrue())

			Expect(set.IsSet[*Person](set1)).To(BeFalse())
			Expect(set.IsSet[*Person](set2)).To(BeFalse())

			Expect(set.IsSet[int](set1)).To(BeFalse())
			Expect(set.IsSet[int](set2)).To(BeFalse())

			Expect(set.IsSet[Person](nil)).To(BeFalse())
			Expect(set.IsSet[int](nil)).To(BeFalse())

			Expect(set.IsSet[Person](p1)).To(BeFalse())

			Expect(set.IsSet[any](set1)).To(BeFalse())
		})

		It("Should union", func() {
			var union = set1.Union(set2)
			Expect(union.Count()).To(Equal(5))

			Expect(union.Has(p1)).To(BeTrue())
			Expect(union.Has(p2)).To(BeTrue())
			Expect(union.Has(p3)).To(BeTrue())
			Expect(union.Has(p4)).To(BeTrue())
			Expect(union.Has(p5)).To(BeTrue())

			var union2 = set2.Union(set1)
			Expect(union2.Count()).To(Equal(5))
			Expect(union2.ToSlice()).To(ConsistOf(union.ToSlice()))
		})

		It("Should intersect", func() {
			var intersection = set1.Intersect(set2)
			Expect(intersection.Count()).To(Equal(3))

			Expect(intersection.Has(p2)).To(BeTrue())
			Expect(intersection.Has(p3)).To(BeTrue())
			Expect(intersection.Has(p4)).To(BeTrue())

			Expect(intersection.Has(p1)).To(BeFalse())
			Expect(intersection.Has(p5)).To(BeFalse())

			var intersection2 = set2.Intersect(set1)
			Expect(intersection2.Count()).To(Equal(3))
			Expect(intersection2.ToSlice()).To(ConsistOf(intersection.ToSlice()))
		})

		It("Should difference", func() {
			var difference = set1.Difference(set2)
			Expect(difference.Count()).To(Equal(1))

			Expect(difference.Has(p1)).To(BeTrue())

			Expect(difference.Has(p2)).To(BeFalse())
			Expect(difference.Has(p3)).To(BeFalse())
			Expect(difference.Has(p4)).To(BeFalse())
			Expect(difference.Has(p5)).To(BeFalse())

			var difference2 = set2.Difference(set1)
			Expect(difference2.Count()).To(Equal(1))
			Expect(difference2.Has(p5)).To(BeTrue())
		})

		It("Should symmetric difference", func() {
			var symmetricDifference = set1.SymmetricDifference(set2)
			Expect(symmetricDifference.Count()).To(Equal(2))

			Expect(symmetricDifference.Has(p1)).To(BeTrue())
			Expect(symmetricDifference.Has(p5)).To(BeTrue())

			Expect(symmetricDifference.Has(p2)).To(BeFalse())
			Expect(symmetricDifference.Has(p3)).To(BeFalse())
			Expect(symmetricDifference.Has(p4)).To(BeFalse())

			var symmetricDifference2 = set2.SymmetricDifference(set1)
			Expect(symmetricDifference2.Count()).To(Equal(2))
			Expect(symmetricDifference2.ToSlice()).To(ConsistOf(symmetricDifference.ToSlice()))
		})

		It("Should map the new Set", func() {
			var result = set1.Map(func(index int, person Person) any {
				return person.Name
			})
			Expect(result.Count()).To(Equal(4))
			Expect(result.ToSlice()).To(ConsistOf("Alice", "Bob", "Charlie", "David"))
		})

		It("Should reduce to a value", func() {
			var result = set1.Reduce(
				func(accumulator any, person Person) any {
					return accumulator.(int) + person.Age
				},
				0,
			)

			Expect(result).To(Equal(20 + 25 + 30 + 35))
		})

		It("Should group by age", func() {
			var result = set1.Union(set2).GroupBy(func(person Person) any {
				if person.Age < 30 {
					return "Under30"
				} else {
					return "Above30"
				}
			})

			Expect(result["Under30"].Count()).To(Equal(2))
			Expect(result["Above30"].Count()).To(Equal(3))
		})
	})
})

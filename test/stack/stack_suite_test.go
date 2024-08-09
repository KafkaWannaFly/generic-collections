package stack_test

import (
	"github.com/KafkaWannaFly/generic-collections/stack"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestStack(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Stack Suite")
}

var _ = Describe("Test Stack", func() {
	Context("Using integers", func() {
		var integerStack *stack.Stack[int]

		BeforeEach(func() {
			integerStack = stack.From(1, 2, 3, 4, 5)

			Expect(integerStack.Count()).To(Equal(5))
		})

		It("Should push an item on top", func() {
			integerStack.Push(0)

			Expect(integerStack.Count()).To(Equal(6))
			Expect(integerStack.Peek()).To(Equal(0))
			Expect(integerStack.GetAt(0)).To(Equal(0))
			Expect(integerStack.Pop()).To(Equal(0))
		})

		It("Should pop an item from the top", func() {
			item := integerStack.Pop()

			Expect(integerStack.Count()).To(Equal(4))
			Expect(item).To(Equal(1))
		})

		It("Should try to pop an item from the top", func() {
			item, ok := integerStack.TryPop()

			Expect(integerStack.Count()).To(Equal(4))
			Expect(item).To(Equal(1))
			Expect(ok).To(BeTrue())

			integerStack.Clear()
			Expect(func() {
				integerStack.Pop()
			}).To(Panic())
		})

		It("Should try to pop an item from the top", func() {
			item, ok := integerStack.TryPop()

			Expect(integerStack.Count()).To(Equal(4))
			Expect(item).To(Equal(1))
			Expect(ok).To(BeTrue())

			integerStack.Clear()
			item, ok = integerStack.TryPop()
			Expect(item).To(Equal(0))
			Expect(ok).To(BeFalse())
		})

		It("Should peek an item from the top", func() {
			item := integerStack.Peek()

			Expect(integerStack.Count()).To(Equal(5))
			Expect(item).To(Equal(1))

			integerStack.Clear()
			Expect(func() {
				integerStack.Peek()
			}).To(Panic())
		})

		It("Should try peek an item from the top", func() {
			item, ok := integerStack.TryPeek()

			Expect(integerStack.Count()).To(Equal(5))
			Expect(item).To(Equal(1))
			Expect(ok).To(BeTrue())

			integerStack.Clear()
			item, ok = integerStack.TryPeek()
			Expect(item).To(Equal(0))
			Expect(ok).To(BeFalse())
		})

		It("Should map and reduce", func() {
			mapped := integerStack.Map(func(item int) any {
				return item * 2
			})

			Expect(mapped.Count()).To(Equal(5))
			Expect(mapped.GetAt(0)).To(Equal(2))

			sum := mapped.Reduce(func(accValue any, item any) any {
				return accValue.(int) + item.(int)
			}, 0)

			Expect(sum).To(Equal(30))
		})

		It("Should group by", func() {
			grouped := integerStack.GroupBy(func(item int) any {
				return item % 2
			})

			Expect(grouped.Count()).To(Equal(2))
			Expect(grouped.Get(0).Count()).To(Equal(2))
			Expect(grouped.Get(1).Count()).To(Equal(3))
		})
	})
})

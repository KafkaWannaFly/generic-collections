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

package queue_test

import (
	"github.com/KafkaWannaFly/generic-collections/queue"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestQueue(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Queue Suite")
}

var _ = Describe("Test Queue", func() {
	var integerQueue *queue.Queue[int]

	BeforeEach(func() {
		integerQueue = queue.From(1, 2, 3, 4, 5)

		Expect(integerQueue.Count()).To(Equal(5))
	})

	It("Should enqueue an item at the end", func() {
		integerQueue.Enqueue(6)

		Expect(integerQueue.Count()).To(Equal(6))
		Expect(integerQueue.GetAt(5)).To(Equal(6))
	})

	It("Should dequeue an item from the front", func() {
		item := integerQueue.Dequeue()

		Expect(integerQueue.Count()).To(Equal(4))
		Expect(item).To(Equal(1))

		integerQueue.Clear()
		Expect(func() {
			integerQueue.Dequeue()
		}).To(Panic())
	})

	It("Should try to dequeue an item from the front", func() {
		item, ok := integerQueue.TryDequeue()

		Expect(integerQueue.Count()).To(Equal(4))
		Expect(item).To(Equal(1))
		Expect(ok).To(BeTrue())

		integerQueue.Clear()
		item, ok = integerQueue.TryDequeue()
		Expect(item).To(Equal(0))
		Expect(ok).To(BeFalse())
	})

	It("Should peek an item from the front", func() {
		item := integerQueue.Peek()

		Expect(integerQueue.Count()).To(Equal(5))
		Expect(item).To(Equal(1))

		integerQueue.Clear()
		Expect(func() {
			integerQueue.Peek()
		}).To(Panic())
	})

	It("Should try peek an item from the front", func() {
		item, ok := integerQueue.TryPeek()

		Expect(integerQueue.Count()).To(Equal(5))
		Expect(item).To(Equal(1))
		Expect(ok).To(BeTrue())

		integerQueue.Clear()
		item, ok = integerQueue.TryPeek()
		Expect(item).To(Equal(0))
		Expect(ok).To(BeFalse())
	})

	It("Should map & reduce the items in the queue", func() {
		result := integerQueue.Map(func(i int, item int) any {
			return item * 2
		})

		Expect(result.Count()).To(Equal(5))
		Expect(result.GetAt(0)).To(Equal(2))
		Expect(result.GetAt(1)).To(Equal(4))
		Expect(result.GetAt(2)).To(Equal(6))
		Expect(result.GetAt(3)).To(Equal(8))
		Expect(result.GetAt(4)).To(Equal(10))

		sum := result.Reduce(func(accVal any, item any) any {
			return item.(int) + accVal.(int)
		}, 0)
		Expect(sum).To(Equal(30))
	})

	It("Should group the items in the queue by the specified key", func() {
		grouped := integerQueue.GroupBy(func(item int) any {
			if item%2 == 0 {
				return "even"
			} else {
				return "odd"
			}
		})

		Expect(grouped.Count()).To(Equal(2))
		Expect(grouped.Get("even").Count()).To(Equal(2))
		Expect(grouped.Get("odd").Count()).To(Equal(3))
	})
})

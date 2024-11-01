package hashmap_test

import (
	"github.com/KafkaWannaFly/generic-collections/hashmap"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestHashmap(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Hashmap Suite")
}

var _ = Describe("Hashmap", func() {
	When("Use basic datatypes as key and value", func() {
		var stringMap *hashmap.HashMap[string, string]
		BeforeEach(func() {
			stringMap = hashmap.Of(map[string]string{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
				"key4": "value4",
				"key5": "value5",
			})

			Expect(stringMap.Count()).To(Equal(5))
		})

		It("Should not add new entry with same key and override value", func() {
			stringMap.Put("key1", "value6")
			stringMap.Put("key2", "value7")
			stringMap.Put("key3", "value8")
			stringMap.Put("key4", "value9")
			stringMap.Put("key5", "value10")

			Expect(stringMap.Count()).To(Equal(5))

			Expect(stringMap.Get("key1")).To(Equal("value6"))
			Expect(stringMap.Get("key2")).To(Equal("value7"))
			Expect(stringMap.Get("key3")).To(Equal("value8"))
			Expect(stringMap.Get("key4")).To(Equal("value9"))
			Expect(stringMap.Get("key5")).To(Equal("value10"))
		})

		It("Should be able to get the value of a key", func() {
			Expect(stringMap.Get("key1")).To(Equal("value1"))
			Expect(stringMap.Get("key2")).To(Equal("value2"))
			Expect(stringMap.Get("key3")).To(Equal("value3"))
			Expect(stringMap.Get("key4")).To(Equal("value4"))
			Expect(stringMap.Get("key5")).To(Equal("value5"))
		})

		It("Should not be able to get the value of a key", func() {
			Expect(stringMap.Get("key6")).To(BeEmpty())
			Expect(stringMap.Get("")).To(BeEmpty())
		})

		It("Should be able to remove a key", func() {
			var v1 = stringMap.Remove("key1")
			Expect(v1).To(Equal("value1"))
			Expect(stringMap.Count()).To(Equal(4))
		})

		It("Should not remove a key", func() {
			var v6 = stringMap.Remove("key6")
			Expect(v6).To(BeEmpty())
			Expect(stringMap.Count()).To(Equal(5))
		})

		It("Should find the key", func() {
			var k1 = stringMap.Find(func(key string, value string) bool {
				return value == "value1"
			})

			Expect(k1).To(Equal("key1"))
		})

		It("Should not find the key", func() {
			var k6 = stringMap.Find(func(key string, value string) bool {
				return value == "value6"
			})

			Expect(k6).To(BeEmpty())
		})

		It("Should be created from list of Entry", func() {
			var entries = []*hashmap.Entry[string, string]{
				{Key: "key6", Value: "value6"},
				{Key: "key7", Value: "value7"},
				{Key: "key8", Value: "value8"},
				{Key: "key9", Value: "value9"},
				{Key: "key10", Value: "value10"},
			}
			var newMap = hashmap.From(entries...)

			Expect(newMap.Count()).To(Equal(5))
		})

		It("Should loop through all entries", func() {
			var count = 0
			stringMap.ForEach(func(key string, value string) {
				count++
			})

			Expect(count).To(Equal(stringMap.Count()))
		})

		It("Should add all entries from another map", func() {
			var newMap = hashmap.New[string, string]()
			newMap.Put("key6", "value6")
			newMap.Put("key7", "value7")
			newMap.Put("key8", "value8")

			stringMap.AddAll(newMap.ToSlice()...)

			Expect(stringMap.Count()).To(Equal(8))
			Expect(stringMap.HasAll(newMap.ToSlice()...)).To(BeTrue())

			Expect(stringMap.HasKey("key6")).To(BeTrue())
			Expect(stringMap.HasKey("key7")).To(BeTrue())
			Expect(stringMap.HasKey("key8")).To(BeTrue())
			Expect(stringMap.HasKey("key9")).To(BeFalse())

			Expect(stringMap.HasAllKey(newMap.Keys())).To(BeTrue())
		})

		It("Should not has all entries from another map", func() {
			var newMap = hashmap.New[string, string]()
			newMap.Put("key6", "value6")
			newMap.Put("key7", "value7")
			newMap.Put("key8", "value8")

			Expect(stringMap.HasAll(newMap.ToSlice()...)).To(BeFalse())
			Expect(newMap.HasAll(stringMap.ToSlice()...)).To(BeFalse())

			Expect(stringMap.HasAllKey(newMap.Keys())).To(BeFalse())
			Expect(newMap.HasAllKey(stringMap.Keys())).To(BeFalse())
		})

		It("Should has any entries from another map", func() {
			var newMap = hashmap.New[string, string]()
			newMap.Put("key5", "value5")
			newMap.Put("key6", "value6")
			newMap.Put("key7", "value7")
			newMap.Put("key8", "value8")

			Expect(stringMap.HasAny(newMap.ToSlice()...)).To(BeTrue())

			var newMap2 = hashmap.From(
				hashmap.NewEntry("key6", "value6"),
			)
			Expect(stringMap.HasAny(newMap2.ToSlice()...)).To(BeFalse())
		})

		It("Should has any key from an array", func() {
			Expect(stringMap.HasAnyKey([]string{"key1", "key6"})).To(BeTrue())

			Expect(stringMap.HasAnyKey([]string{"key6", "key7"})).To(BeFalse())
		})

		It("Should be able to get all keys", func() {
			var keys = stringMap.Keys()
			Expect(len(keys)).To(Equal(5))
			Expect(keys).To(ContainElements("key1", "key2", "key3", "key4", "key5"))
		})

		It("Should be able to get all values", func() {
			var values = stringMap.Values()
			Expect(len(values)).To(Equal(5))
			Expect(values).To(ContainElements("value1", "value2", "value3", "value4", "value5"))
		})

		It("Should be able to get all entries", func() {
			var entries = stringMap.Entries()
			Expect(len(entries)).To(Equal(5))

			Expect(entries).To(
				ContainElements(
					hashmap.NewEntry("key1", "value1"),
					hashmap.NewEntry("key2", "value2"),
					hashmap.NewEntry("key3", "value3"),
					hashmap.NewEntry("key4", "value4"),
					hashmap.NewEntry("key5", "value5"),
				),
			)
		})

		It("Should be able to clear the map", func() {
			stringMap.Clear()
			Expect(stringMap.Count()).To(Equal(0))
			Expect(stringMap.IsEmpty()).To(BeTrue())
		})

		It("Should be able to filter the map", func() {
			var filtered = stringMap.Filter(func(key string, value string) bool {
				return value == "value1" || value == "value2"
			})
			Expect(filtered.Count()).To(Equal(2))

			Expect(filtered.Get("key1")).To(Equal("value1"))
			Expect(filtered.Get("key2")).To(Equal("value2"))

			Expect(filtered.HasKey("key3")).To(BeFalse())
		})

		It("Should clone the map", func() {
			var cloned = stringMap.Clone()

			Expect(cloned.Count()).To(Equal(stringMap.Count()))
			Expect(cloned.HasAll(stringMap.Entries()...)).To(BeTrue())
			Expect(cloned).NotTo(BeIdenticalTo(stringMap))

			cloned.ForEach(func(key string, value string) {
				Expect(stringMap.Get(key)).To(Equal(value))
			})
		})

		It("Should clone entry", func() {
			var entry = hashmap.NewEntry("key1", "value1")
			var cloned = entry.Clone()

			Expect(cloned).To(Equal(entry))
			Expect(cloned.Key).To(Equal(entry.Key))
			Expect(cloned.Value).To(Equal(entry.Value))
			Expect(cloned.Equals(entry)).To(BeTrue())
		})
	})
})

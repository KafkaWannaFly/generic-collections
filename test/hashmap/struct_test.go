package hashmap

import (
	"fmt"
	"generic-collections/hashmap"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type City struct {
	Name       string
	CountyName string
}

func (c City) GetHashCode() string {
	return fmt.Sprintf("%s-%s", c.Name, c.CountyName)
}

type District struct {
	Name         string
	Population   int
	SubDistricts []string
}

func (receiver District) GetHashCode() string {
	return fmt.Sprintf("%s-%d-%v", receiver.Name, receiver.Population, receiver.SubDistricts)
}

var _ = Describe("Hashmap Struct Test", func() {
	When("Use struct as key and value", func() {
		var cityMap *hashmap.HashMap[City, District]

		var K1 = City{"Istanbul", "Turkey"}
		var V1 = District{"Kadikoy", 100000, []string{"Moda", "Goztepe"}}

		var K2 = City{"Ankara", "Turkey"}
		var V2 = District{"Cankaya", 50000, []string{"Kizilay", "Bahcelievler"}}

		var K3 = City{"Izmir", "Turkey"}
		var V3 = District{"Bornova", 75000, []string{"Ege", "Evka"}}

		var K4 = City{"Paris", "France"}
		var V4 = District{"Le Marais", 25000, []string{"Saint-Paul", "Saint-Gervais"}}

		var K5 = City{"London", "United Kingdom"}
		var V5 = District{"Soho", 30000, []string{"Mayfair", "Fitzrovia"}}

		BeforeEach(func() {
			cityMap = hashmap.New[City, District]()

			cityMap.Put(K1, V1).Put(K2, V2).Put(K3, V3).Put(K4, V4).Put(K5, V5)

			Expect(cityMap.Count()).To(Equal(5))
		})

		It("Should check data type", func() {
			Expect(hashmap.IsHashMap[City, District](cityMap)).To(BeTrue())

			Expect(hashmap.IsHashMap[string, string](cityMap)).To(BeFalse())
			Expect(hashmap.IsHashMap[*City, *District](cityMap)).To(BeFalse())
			Expect(hashmap.IsHashMap[City, District](nil)).To(BeFalse())
			Expect(hashmap.IsHashMap[City, District](K1)).To(BeFalse())
		})

		It("Should not add new entry with same key and override value", func() {
			var c1 = City{"Istanbul", "Turkey"}
			var v1 = District{"Besiktas", 200000, []string{"Levent", "Etiler"}}
			cityMap.Set(c1, v1)

			var c2 = City{"Ankara", "Turkey"}
			var v2 = District{"Mamak", 150000, []string{"Sincan", "Cayyolu"}}
			cityMap.Set(c2, v2)

			var c3 = City{"Izmir", "Turkey"}
			var v3 = District{"Konak", 100000, []string{"Alsancak", "Karsiyaka"}}
			cityMap.Set(c3, v3)

			var c4 = City{"Paris", "France"}
			var v4 = District{"Montmartre", 50000, []string{"Pigalle", "Clignancourt"}}
			cityMap.Set(c4, v4)

			var c5 = City{"London", "United Kingdom"}
			var v5 = District{"Covent Garden", 40000, []string{"Holborn", "Strand"}}
			cityMap.Set(c5, v5)

			Expect(cityMap.Count()).To(Equal(5))

			Expect(cityMap.Get(c1)).To(Equal(v1))
			Expect(cityMap.Get(c2)).To(Equal(v2))
			Expect(cityMap.Get(c3)).To(Equal(v3))
			Expect(cityMap.Get(c4)).To(Equal(v4))
			Expect(cityMap.Get(c5)).To(Equal(v5))
		})

		It("Should be able to get the value of a key", func() {
			Expect(cityMap.Get(K1)).To(Equal(V1))
			Expect(cityMap.Get(K2)).To(Equal(V2))
			Expect(cityMap.Get(K3)).To(Equal(V3))
			Expect(cityMap.Get(K4)).To(Equal(V4))
			Expect(cityMap.Get(K5)).To(Equal(V5))
		})

		It("Should not be able to get the value of a key", func() {
			var c6 = City{"Berlin", "Germany"}
			Expect(cityMap.Get(c6)).To(Equal(District{}))
		})

		It("Should be able to remove a key", func() {
			var v1 = cityMap.Remove(K1)
			Expect(v1).To(Equal(V1))
			Expect(cityMap.Count()).To(Equal(4))

			var v2 = cityMap.Remove(K2)
			Expect(v2).To(Equal(V2))
			Expect(cityMap.Count()).To(Equal(3))

			var v3 = cityMap.Remove(K3)
			Expect(v3).To(Equal(V3))
			Expect(cityMap.Count()).To(Equal(2))

			var v4 = cityMap.Remove(K4)
			Expect(v4).To(Equal(V4))
			Expect(cityMap.Count()).To(Equal(1))

			var v5 = cityMap.Remove(K5)
			Expect(v5).To(Equal(V5))
			Expect(cityMap.Count()).To(Equal(0))
		})

		It("Should not remove a key", func() {
			var c6 = City{"Berlin", "Germany"}
			var v6 = cityMap.Remove(c6)
			Expect(v6).To(Equal(District{}))
			Expect(cityMap.Count()).To(Equal(5))
		})

		It("Should be created from list of Entry", func() {
			var entries = []hashmap.Entry[City, District]{
				hashmap.NewEntry(K1, V1),
				hashmap.NewEntry(K2, V2),
				hashmap.NewEntry(K3, V3),
				hashmap.NewEntry(K4, V4),
				hashmap.NewEntry(K5, V5),
			}

			var newCityMap = hashmap.From(entries...)
			Expect(newCityMap.Count()).To(Equal(5))

			Expect(newCityMap.Get(K1)).To(Equal(V1))
			Expect(newCityMap.Get(K2)).To(Equal(V2))
			Expect(newCityMap.Get(K3)).To(Equal(V3))
			Expect(newCityMap.Get(K4)).To(Equal(V4))
			Expect(newCityMap.Get(K5)).To(Equal(V5))
		})

		It("Should loop through all entries", func() {
			var count = 0
			cityMap.ForEach(func(_ int, entry hashmap.Entry[City, District]) {
				count++
			})

			Expect(count).To(Equal(cityMap.Count()))
		})

		It("Should add all entries from another map", func() {
			var newCityMap = hashmap.New[City, District]()
			newCityMap.Put(
				City{"Berlin", "Germany"},
				District{"Mitte", 150000, []string{"Tiergarten", "Wedding"}},
			).Put(
				City{"Madrid", "Spain"},
				District{"Sol", 200000, []string{"Gran Via", "Chueca"}},
			).Put(
				K1,
				V1,
			)

			cityMap.AddAll(newCityMap)

			Expect(cityMap.Count()).To(Equal(7))
		})

		It("Should not has all entries from another map", func() {
			var newCityMap = hashmap.New[City, District]()
			newCityMap.Put(
				City{"Berlin", "Germany"},
				District{"Mitte", 150000, []string{"Tiergarten", "Wedding"}},
			).Put(
				City{"Madrid", "Spain"},
				District{"Sol", 200000, []string{"Gran Via", "Chueca"}},
			).Put(K1, V1)

			Expect(cityMap.HasAll(newCityMap)).To(BeFalse())
		})

		It("Should be able to get all keys", func() {
			var keys = cityMap.GetKeys()
			Expect(keys).To(ContainElements(K1, K2, K3, K4, K5))
		})

		It("Should be able to get all values", func() {
			var values = cityMap.GetValues()
			Expect(values).To(ContainElements(V1, V2, V3, V4, V5))
		})

		It("Should be able to get all entries", func() {
			var entries = cityMap.GetEntries()
			Expect(entries).To(ContainElements(
				hashmap.NewEntry(K1, V1),
				hashmap.NewEntry(K2, V2),
				hashmap.NewEntry(K3, V3),
				hashmap.NewEntry(K4, V4),
				hashmap.NewEntry(K5, V5),
			))
		})

		It("Should be able to clear all entries", func() {
			cityMap.Clear()
			Expect(cityMap.Count()).To(Equal(0))

			Expect(cityMap.GetKeys()).To(BeEmpty())
			Expect(cityMap.GetValues()).To(BeEmpty())
		})

		It("Should be able to filter entries", func() {
			var filtered = cityMap.Filter(func(entry hashmap.Entry[City, District]) bool {
				return entry.Key.Name == "Istanbul"
			}).(*hashmap.HashMap[City, District])

			Expect(filtered.Count()).To(Equal(1))
			Expect(filtered.GetKeys()).To(ContainElements(K1))
			Expect(filtered.GetValues()).To(ContainElements(V1))
		})
	})
})

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

type Book struct {
	Title         string
	Author        string
	CurrentPage   int
	Pages         []string
	PublishedYear int
	Price         float64
}

func (receiver Book) Compare(book Book) int {
	if receiver.Title == book.Title {
		return 0
	} else if receiver.Title > book.Title {
		return 1
	} else {
		return -1
	}
}

func (receiver Book) GetHashCode() string {
	return receiver.Title
}

var _ = Describe("Test List implements ICollection", func() {

	Context("Using integer", func() {
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
			integerList.Set(0, 11).Set(1, 12).Set(2, 13).Set(3, 14).Set(4, 15).Set(5, 16).Set(6, 17).Set(7, 18).Set(8, 19).Set(9, 20)

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
			Expect(integerList.Has(5)).To(BeTrue())
			Expect(integerList.Has(15)).To(BeFalse())
		})

		It("Should check if contains all elements", func() {
			Expect(integerList.HasAll(list.From(1, 2, 3, 4, 5))).To(BeTrue())
			Expect(integerList.HasAll(list.From(1, 2, 3, 4, 15))).To(BeFalse())
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

		It("Should add elements", func() {
			stringList.Add("Lemon").Add("Mango").Add("Nectarine").Add("Orange").Add("Papaya").Add("Quince").Add("Raspberry").Add("Strawberry").Add("Tangerine").Add("Ugli")

			Expect(stringList.Count()).To(Equal(20))
		})

		It("Should add all elements", func() {
			stringList.AddAll(list.From("Lemon", "Mango", "Nectarine", "Orange", "Papaya", "Quince", "Raspberry", "Strawberry", "Tangerine", "Ugli"))

			Expect(stringList.Count()).To(Equal(20))
		})

		It("Should clear all elements", func() {
			stringList.Clear()

			Expect(stringList.Count()).To(Equal(0))
			Expect(stringList.IsEmpty()).To(BeTrue())
		})

		It("Should contain an element", func() {
			Expect(stringList.Has("Apple")).To(BeTrue())
			Expect(stringList.Has("Elderberry")).To(BeTrue())
			Expect(stringList.Has("Devil Fruit")).To(BeFalse())
		})

		It("Should contain all elements", func() {
			Expect(stringList.HasAll(list.From("Apple", "Banana", "Cherry", "Dates", "Elderberry"))).To(BeTrue())
			Expect(stringList.HasAll(list.From("Apple", "Banana", "Cherry", "Dates", "Devil Fruit"))).To(BeFalse())
		})

		It("Should filter elements", func() {
			filtered := stringList.Filter(func(element string) bool {
				return element != "Apple"
			})

			Expect(filtered.Count()).To(Equal(9))
			Expect(filtered.Has("Apple")).To(BeFalse())
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
			stringList.Set(0, "Lemon").Set(1, "Mango").Set(2, "Nectarine").Set(3, "Orange").Set(4, "Papaya").Set(5, "Quince")

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

		It("Should iterate over elements", func() {
			var sum int
			stringList.ForEach(func(_ int, element string) {
				sum += 1
			})

			Expect(sum).To(Equal(10))
		})
	})

	Context("Using struct", func() {
		var bookList *list.List[Book]

		BeforeEach(func() {
			bookList = list.New[Book]()
			bookList.Add(
				Book{
					Title:       "The Alchemist",
					Author:      "Paulo Coelho",
					CurrentPage: 0,
					Pages: []string{
						"Prologue",
						"Part One",
						"Part Two",
						"Part Three",
					},
					PublishedYear: 1988,
					Price:         10.99,
				}).Add(
				Book{
					Title:       "The Little Prince",
					Author:      "Antoine de Saint-Exupéry",
					CurrentPage: 0,
					Pages: []string{
						"Chapter 1",
						"Chapter 2",
						"Chapter 3",
						"Chapter 4",
					},
					PublishedYear: 1943,
					Price:         9.99,
				}).Add(
				Book{
					Title:       "The Catcher in the Rye",
					Author:      "J. D. Salinger",
					CurrentPage: 0,
					Pages: []string{
						"Chapter 1",
						"Chapter 2",
						"Chapter 3",
						"Chapter 4",
					},
					PublishedYear: 1951,
					Price:         11.99,
				})

			Expect(bookList.Count()).To(Equal(3))
		})

		It("Should add elements", func() {
			bookList.Add(
				Book{
					Title: "The Great Gatsby",
				})

			Expect(bookList.Count()).To(Equal(4))
		})

		It("Should add all elements", func() {
			bookList.AddAll(list.From(
				Book{
					Title: "The Great Gatsby",
				},
				Book{
					Title: "To Kill a Mockingbird",
				},
				Book{
					Title: "1984",
				},
				Book{
					Title: "Animal Farm",
				},
				Book{
					Title: "Brave New World",
				},
			))

			Expect(bookList.Count()).To(Equal(8))
		})

		It("Should clear all elements", func() {
			bookList.Clear()

			Expect(bookList.Count()).To(Equal(0))
			Expect(bookList.IsEmpty()).To(BeTrue())
		})

		It("Should contain an element", func() {
			Expect(bookList.Has(Book{
				Title: "The Alchemist",
			})).To(BeTrue())

			Expect(bookList.Has(Book{
				Title: "The Great Gatsby",
			})).To(BeFalse())
		})

		It("Should contain all elements", func() {
			Expect(bookList.HasAll(list.From(
				Book{
					Title: "The Alchemist",
				},
				Book{
					Title: "The Little Prince",
				},
				Book{
					Title: "The Catcher in the Rye",
				},
			))).To(BeTrue())

			Expect(bookList.HasAll(list.From(
				Book{
					Title: "The Alchemist",
				},
				Book{
					Title: "The Great Gatsby",
				},
			))).To(BeFalse())
		})

		It("Should filter elements", func() {
			filtered := bookList.Filter(func(book Book) bool {
				return book.PublishedYear > 1950
			})

			Expect(filtered.Count()).To(Equal(2))
		})

		It("Should get elements", func() {
			Expect(bookList.Get(0).Title).To(Equal("The Alchemist"))
			Expect(bookList.Get(1).Title).To(Equal("The Little Prince"))
			Expect(bookList.Get(2).Title).To(Equal("The Catcher in the Rye"))
		})

		It("Should set elements", func() {
			bookList.Set(0, Book{
				Title: "The Great Gatsby",
			}).Set(1, Book{
				Title: "To Kill a Mockingbird",
			}).Set(2, Book{
				Title: "1984",
			})

			Expect(bookList.Get(0).Title).To(Equal("The Great Gatsby"))
			Expect(bookList.Get(1).Title).To(Equal("To Kill a Mockingbird"))
			Expect(bookList.Get(2).Title).To(Equal("1984"))
		})

		It("Should convert to slice", func() {
			var bookSlice = bookList.ToSlice()
			Expect(bookSlice[0].Title).To(Equal("The Alchemist"))
			Expect(bookSlice[1].Title).To(Equal("The Little Prince"))
			Expect(bookSlice[2].Title).To(Equal("The Catcher in the Rye"))

			Expect(len(bookSlice)).To(Equal(3))
		})

		It("Should iterate over elements", func() {
			var sum int
			bookList.ForEach(func(index int, book Book) {
				sum += 1

				Expect(bookList.Get(index).Title).To(Equal(book.Title))
			})

			Expect(sum).To(Equal(3))
		})
	})
})

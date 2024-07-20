package collection_test

import (
	"generic-collections/interfaces"
	"generic-collections/linkedlist"
	"generic-collections/list"
	"generic-collections/set"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCollection(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Collection Suite")
}

var _ = Describe("Test Collection", func() {

	When("Using integer", func() {
		var integerList = list.From(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		Context("For List", integerTests(integerList))

		var integerSet = set.From(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		Context("For Set", integerTests(integerSet))

		var integerLinkedList = linkedlist.From(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		Context("For LinkedList", integerTests(integerLinkedList))
	})

	When("Using string", func() {
		var stringList = list.From(
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
		Context("For List", stringTests(stringList))

		var stringSet = set.From(stringList.ToSlice()...)
		Context("For Set", stringTests(stringSet))

		var stringLinkedList = linkedlist.From(stringList.ToSlice()...)
		Context("For LinkedList", stringTests(stringLinkedList))
	})

	When("Using struct", func() {
		var bookList = list.New[Book]()
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
			},
		).Add(
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
			},
		).Add(
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
			},
		)
		Context("For List", structTests(bookList))

		var bookSet = set.From(bookList.ToSlice()...)
		Context("For Set", structTests(bookSet))

		var bookLinkedList = linkedlist.From(bookList.ToSlice()...)
		Context("For LinkedList", structTests(bookLinkedList))
	})

	When("Using pointer", func() {
		var studentList = list.New[*Student]()
		studentList.Add(
			&Student{
				Name: "Alice",
				Age:  20,
				GPA:  3.5,
			}).Add(
			&Student{
				Name: "Bob",
				Age:  21,
				GPA:  3.6,
			}).Add(
			&Student{
				Name: "Charlie",
				Age:  22,
				GPA:  3.7,
			})
		Context("For List", pointerTests(studentList))

		var studentSet = set.From(studentList.ToSlice()...)
		Context("For Set", pointerTests(studentSet))

		var studentLinkedList = linkedlist.From(studentList.ToSlice()...)
		Context("For LinkedList", pointerTests(studentLinkedList))
	})
})

func integerTests(collection interfaces.ICollection[int]) func() {

	return func() {
		var integerCollection interfaces.ICollection[int]

		BeforeEach(func() {
			integerCollection = collection.Clone()
		})

		It("Should add an element", func() {
			integerCollection.Add(11)

			Expect(integerCollection.Count()).To(Equal(11))
		})

		It("Should add all elements", func() {
			integerCollection.AddAll(list.From(11, 12, 13, 14, 15))

			Expect(integerCollection.Count()).To(Equal(15))
		})

		It("Should clear all elements", func() {
			integerCollection.Clear()

			Expect(integerCollection.Count()).To(Equal(0))
			Expect(integerCollection.IsEmpty(), BeTrue())
		})

		It("Should check if contains an element", func() {
			Expect(integerCollection.Has(5)).To(BeTrue())
			Expect(integerCollection.Has(15)).To(BeFalse())
		})

		It("Should check if contains all elements", func() {
			Expect(integerCollection.HasAll(list.From(1, 2, 3, 4, 5))).To(BeTrue())
			Expect(integerCollection.HasAll(list.From(1, 2, 3, 4, 15))).To(BeFalse())
		})

		It("Should check if contains any element", func() {
			Expect(integerCollection.HasAny(list.From(1, 2, 3, 4, 5, 20))).To(BeTrue())
			Expect(integerCollection.HasAny(list.From(11, 12, 13, 14, 15))).To(BeFalse())
		})

		It("Should filter elements", func() {
			filtered := integerCollection.Filter(func(element int) bool {
				return element > 5
			})

			Expect(filtered.Count()).To(Equal(5))
		})

		It("Should iterate over elements", func() {
			var sum int
			integerCollection.ForEach(func(_ int, element int) {
				sum += element
			})

			Expect(sum).To(Equal(55))
		})
	}
}

func stringTests(collection interfaces.ICollection[string]) func() {
	return func() {
		var stringCollection interfaces.ICollection[string]

		BeforeEach(func() {
			stringCollection = collection.Clone()

			Expect(stringCollection.Count()).To(Equal(10))
		})

		It("Should add elements", func() {
			stringCollection.Add("Lemon").Add("Mango").Add("Nectarine").Add("Orange").Add("Papaya").Add("Quince").Add("Raspberry").Add("Strawberry").Add("Tangerine").Add("Ugli")

			Expect(stringCollection.Count()).To(Equal(20))
		})

		It("Should add all elements", func() {
			stringCollection.AddAll(list.From("Lemon", "Mango", "Nectarine", "Orange", "Papaya", "Quince", "Raspberry", "Strawberry", "Tangerine", "Ugli"))

			Expect(stringCollection.Count()).To(Equal(20))
		})

		It("Should clear all elements", func() {
			stringCollection.Clear()

			Expect(stringCollection.Count()).To(Equal(0))
			Expect(stringCollection.IsEmpty()).To(BeTrue())
		})

		It("Should contain an element", func() {
			Expect(stringCollection.Has("Apple")).To(BeTrue())
			Expect(stringCollection.Has("Elderberry")).To(BeTrue())
			Expect(stringCollection.Has("Devil Fruit")).To(BeFalse())
		})

		It("Should contain all elements", func() {
			Expect(stringCollection.HasAll(list.From("Apple", "Banana", "Cherry", "Dates", "Elderberry"))).To(BeTrue())
			Expect(stringCollection.HasAll(list.From("Apple", "Banana", "Cherry", "Dates", "Devil Fruit"))).To(BeFalse())
		})

		It("Should contain any element", func() {
			Expect(stringCollection.HasAny(list.From("Apple", "Banana", "Cherry", "Dates", "Elderberry", "Devil Fruit"))).To(BeTrue())
			Expect(stringCollection.HasAny(list.From("Devil Fruit", "Elderberry", "Fig", "Grape", "Honeydew", "Jackfruit"))).To(BeTrue())
			Expect(stringCollection.HasAny(list.From("Devil Fruit"))).To(BeFalse())
		})

		It("Should filter elements", func() {
			filtered := stringCollection.Filter(func(element string) bool {
				return element != "Apple"
			})

			Expect(filtered.Count()).To(Equal(9))
			Expect(filtered.Has("Apple")).To(BeFalse())
		})

		It("Should iterate over elements", func() {
			var sum int
			stringCollection.ForEach(func(_ int, element string) {
				sum += 1
			})

			Expect(sum).To(Equal(10))
		})
	}
}

// region Structs Definition
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

func (receiver Book) HashCode() string {
	return receiver.Title
}

type Student struct {
	Name string
	Age  int
	GPA  float64
}

func (receiver *Student) Compare(item *Student) int {
	if receiver.Name == item.Name {
		return 0
	} else if receiver.Name > item.Name {
		return 1
	} else {
		return -1
	}
}

func (receiver *Student) HashCode() string {
	return receiver.Name
}

func (receiver *Student) SetName(name string) {
	receiver.Name = name
}

// endregion

func structTests(collection interfaces.ICollection[Book]) func() {
	return func() {
		var bookList interfaces.ICollection[Book]

		BeforeEach(func() {
			bookList = collection.Clone()

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

		It("Should contain any element", func() {
			Expect(bookList.HasAny(list.From(
				Book{
					Title: "The Alchemist",
				},
				Book{
					Title: "The Great Gatsby",
				},
				Book{
					Title: "The Catcher in the Rye",
				},
			))).To(BeTrue())

			Expect(bookList.HasAny(list.From(
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

		It("Should iterate over elements", func() {
			var sum int
			bookList.ForEach(func(index int, book Book) {
				sum += 1
			})

			Expect(sum).To(Equal(3))
		})
	}
}

func pointerTests(collection interfaces.ICollection[*Student]) func() {
	return func() {
		var studentList interfaces.ICollection[*Student]

		BeforeEach(func() {
			studentList = collection.Clone()

			Expect(studentList.Count()).To(Equal(3))
		})

		It("Should add elements", func() {
			Expect(studentList.Count()).To(Equal(3))

			studentList.Add(
				&Student{
					Name: "David",
				})

			Expect(studentList.Count()).To(Equal(4))
		})

		It("Should add all elements", func() {
			studentList.AddAll(list.From(
				&Student{
					Name: "David",
				},
				&Student{
					Name: "Eve",
				},
				&Student{
					Name: "Frank",
				},
				&Student{
					Name: "Grace",
				},
				&Student{
					Name: "Henry",
				},
			))

			Expect(studentList.Count()).To(Equal(8))
		})

		It("Should clear all elements", func() {
			Expect(studentList.Count()).To(Equal(3))
			Expect(studentList.IsEmpty()).To(BeFalse())

			studentList.Clear()

			Expect(studentList.Count()).To(Equal(0))
			Expect(studentList.IsEmpty()).To(BeTrue())
		})

		It("Should contain an element", func() {
			Expect(studentList.Has(&Student{
				Name: "Alice",
			})).To(BeTrue())

			Expect(studentList.Has(&Student{
				Name: "David",
			})).To(BeFalse())
		})

		It("Should contain all elements", func() {
			Expect(studentList.HasAll(list.From(
				&Student{
					Name: "Alice",
				},
				&Student{
					Name: "Bob",
				},
				&Student{
					Name: "Charlie",
				},
			))).To(BeTrue())

			Expect(studentList.HasAll(list.From(
				&Student{
					Name: "Alice",
				},
				&Student{
					Name: "David",
				},
			))).To(BeFalse())
		})

		It("Should filter elements", func() {
			filtered := studentList.Filter(func(student *Student) bool {
				return student.Age > 20
			})

			Expect(filtered.Count()).To(Equal(2))
		})

		It("Should iterate over elements", func() {
			var sum int
			studentList.ForEach(func(index int, student *Student) {
				sum += 1
			})

			Expect(sum).To(Equal(3))
		})
	}
}

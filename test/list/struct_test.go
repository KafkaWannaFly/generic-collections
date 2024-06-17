package list

import (
	"generic-collections/list"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

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

func (receiver *Student) GetHashCode() string {
	return receiver.Name
}

func (receiver *Student) SetName(name string) {
	receiver.Name = name
}

var _ = Describe("Test List implements ICollection", func() {
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
			})
			bookList.Set(1, Book{
				Title: "To Kill a Mockingbird",
			})
			bookList.Set(2, Book{
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

	Context("Using pointer", func() {
		var studentList *list.List[*Student]

		BeforeEach(func() {
			studentList = list.New[*Student]()
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

		It("Should get elements", func() {
			Expect(studentList.Get(0).Name).To(Equal("Alice"))
			Expect(studentList.Get(1).Name).To(Equal("Bob"))
			Expect(studentList.Get(2).Name).To(Equal("Charlie"))
		})

		It("Should set elements", func() {
			studentList.Set(0, &Student{
				Name: "David",
			})
			studentList.Set(1, &Student{
				Name: "Eve",
			})
			studentList.Set(2, &Student{
				Name: "Frank",
			})

			Expect(studentList.Get(0).Name).To(Equal("David"))
			Expect(studentList.Get(1).Name).To(Equal("Eve"))
			Expect(studentList.Get(2).Name).To(Equal("Frank"))
		})

		It("Should convert to slice", func() {
			var studentSlice = studentList.ToSlice()
			Expect(studentSlice[0].Name).To(Equal("Alice"))
			Expect(studentSlice[1].Name).To(Equal("Bob"))
			Expect(studentSlice[2].Name).To(Equal("Charlie"))

			Expect(len(studentSlice)).To(Equal(3))
		})

		It("Should iterate over elements", func() {
			var sum int
			studentList.ForEach(func(index int, student *Student) {
				sum += 1

				Expect(studentList.Get(index).Name).To(Equal(student.Name))
			})

			Expect(sum).To(Equal(3))
		})

		It("Should update elements", func() {
			var student = studentList.Get(0)
			student.SetName("Kafka")

			Expect(studentList.Get(0).Name).To(Equal("Kafka"))
		})
	})
})

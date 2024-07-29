package list

import (
	"generic-collections/list"
	"generic-collections/utils"
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
			Expect(list.IsList[Book](bookList)).To(BeTrue())
		})

		It("Should get elements", func() {
			Expect(bookList.GetAt(0).Title).To(Equal("The Alchemist"))
			Expect(bookList.GetAt(1).Title).To(Equal("The Little Prince"))
			Expect(bookList.GetAt(2).Title).To(Equal("The Catcher in the Rye"))
		})

		It("Should set elements", func() {
			bookList.SetAt(0, Book{
				Title: "The Great Gatsby",
			})
			bookList.SetAt(1, Book{
				Title: "To Kill a Mockingbird",
			})
			bookList.SetAt(2, Book{
				Title: "1984",
			})

			Expect(bookList.GetAt(0).Title).To(Equal("The Great Gatsby"))
			Expect(bookList.GetAt(1).Title).To(Equal("To Kill a Mockingbird"))
			Expect(bookList.GetAt(2).Title).To(Equal("1984"))
		})

		It("Should remove elements", func() {
			var book0 = bookList.RemoveAt(0)
			Expect(book0.Title).To(Equal("The Alchemist"))

			var book1 = bookList.RemoveAt(0)
			Expect(book1.Title).To(Equal("The Little Prince"))

			var book2 = bookList.RemoveAt(0)
			Expect(book2.Title).To(Equal("The Catcher in the Rye"))

			Expect(bookList.Count()).To(Equal(0))
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

				Expect(bookList.GetAt(index).Title).To(Equal(book.Title))
			})

			Expect(sum).To(Equal(3))
		})

		It("Should find elements", func() {
			var index = bookList.FindFirst(func(_ int, book Book) bool {
				return book.Title == "The Little Prince"
			})

			Expect(index).To(Equal(1))
			Expect(bookList.GetAt(index).Title).To(Equal("The Little Prince"))
		})

		It("Should not find elements", func() {
			var index = bookList.FindFirst(func(_ int, book Book) bool {
				return book.Title == "The Great Gatsby"
			})

			Expect(index).To(Equal(-1))
		})

		It("Should map to another type", func() {
			var titleList = bookList.Map(func(index int, book Book) any {
				return book.Title
			})

			Expect(titleList.GetAt(0)).To(Equal("The Alchemist"))
			Expect(titleList.GetAt(1)).To(Equal("The Little Prince"))
			Expect(titleList.GetAt(2)).To(Equal("The Catcher in the Rye"))

			// Type of titleList is *List[any] instead of *List[string]
			// Because List can't handle 2nd type parameter
			Expect(list.IsList[any](titleList)).To(BeTrue())

			var titleList2 = list.Map(bookList, func(index int, book Book) string {
				return book.Title
			})
			// Type of titleList2 is *List[string]
			Expect(list.IsList[string](titleList2)).To(BeTrue())

			Expect(titleList.ToSlice()).To(ContainElements(titleList2.ToSlice()))
		})

		It("Should reduce elements", func() {
			var total = bookList.Reduce(func(result any, book Book) any {
				return result.(float64) + book.Price
			}, 0.0).(float64)

			var total2 = list.Reduce(bookList, func(result float64, book Book) float64 {
				return result + book.Price
			}, 0)

			Expect(total).To(Equal(32.97))
			Expect(total).To(Equal(total2))
		})

		It("Should map then reduce elements", func() {
			var total = bookList.Map(func(index int, book Book) any {
				return book.Price
			}).Reduce(func(result any, price any) any {
				return result.(float64) + price.(float64)
			}, 0.0).(float64)

			Expect(total).To(Equal(32.97))
		})

		It("Should group elements", func() {
			var groups = bookList.GroupBy(func(book Book) any {
				return book.PublishedYear
			})

			Expect(groups.Count()).To(Equal(3))

			var group1943 = groups.Get(1943)
			Expect(group1943.Count()).To(Equal(1))
			Expect(group1943.GetAt(0).Title).To(Equal("The Little Prince"))

			var group1951 = groups.Get(1951)
			Expect(group1951.Count()).To(Equal(1))
			Expect(group1951.GetAt(0).Title).To(Equal("The Catcher in the Rye"))

			var group1988 = groups.Get(1988)
			Expect(group1988.Count()).To(Equal(1))
			Expect(group1988.GetAt(0).Title).To(Equal("The Alchemist"))
		})

		It("Should try get elements in range", func() {
			var book, ok = bookList.TryGetAt(0)
			Expect(ok).To(BeTrue())
			Expect(book.Title).To(Equal("The Alchemist"))

			book, ok = bookList.TryGetAt(1)
			Expect(ok).To(BeTrue())
			Expect(book.Title).To(Equal("The Little Prince"))

			book, ok = bookList.TryGetAt(2)
			Expect(ok).To(BeTrue())
			Expect(book.Title).To(Equal("The Catcher in the Rye"))
		})

		It("Should try get elements out of range", func() {
			var book, ok = bookList.TryGetAt(-1)
			Expect(ok).To(BeFalse())
			Expect(book).To(Equal(utils.DefaultValue[Book]()))

			book, ok = bookList.TryGetAt(3)
			Expect(ok).To(BeFalse())
			var defaultBook Book
			Expect(book).To(Equal(defaultBook))
		})

		It("Should try set elements in range", func() {
			var ok = bookList.TrySetAt(0, Book{
				Title: "The Great Gatsby",
			})
			Expect(ok).To(BeTrue())
			Expect(bookList.GetAt(0).Title).To(Equal("The Great Gatsby"))

			ok = bookList.TrySetAt(1, Book{
				Title: "To Kill a Mockingbird",
			})
			Expect(ok).To(BeTrue())
			Expect(bookList.GetAt(1).Title).To(Equal("To Kill a Mockingbird"))

			ok = bookList.TrySetAt(2, Book{
				Title: "1984",
			})
			Expect(ok).To(BeTrue())
			Expect(bookList.GetAt(2).Title).To(Equal("1984"))
		})

		It("Should try set elements out of range", func() {
			var ok = bookList.TrySetAt(-1, Book{
				Title: "The Great Gatsby",
			})
			Expect(ok).To(BeFalse())

			ok = bookList.TrySetAt(3, Book{
				Title: "To Kill a Mockingbird",
			})
			Expect(ok).To(BeFalse())
		})

		It("Should try remove elements in range", func() {
			var book, ok = bookList.TryRemoveAt(0)
			Expect(ok).To(BeTrue())
			Expect(book.Title).To(Equal("The Alchemist"))

			book, ok = bookList.TryRemoveAt(0)
			Expect(ok).To(BeTrue())
			Expect(book.Title).To(Equal("The Little Prince"))

			book, ok = bookList.TryRemoveAt(0)
			Expect(ok).To(BeTrue())
			Expect(book.Title).To(Equal("The Catcher in the Rye"))

			Expect(bookList.Count()).To(Equal(0))
		})

		It("Should try remove elements out of range", func() {
			var book, ok = bookList.TryRemoveAt(-1)
			Expect(ok).To(BeFalse())
			Expect(book).To(Equal(utils.DefaultValue[Book]()))

			book, ok = bookList.TryRemoveAt(11)
			Expect(ok).To(BeFalse())
			Expect(book).To(Equal(utils.DefaultValue[Book]()))
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

		It("Should assert correct data type", func() {
			Expect(list.IsList[*Student](studentList)).To(BeTrue())

			Expect(list.IsList[*Student](nil)).To(BeFalse())

			Expect(list.IsList[Student](studentList)).To(BeFalse())
			Expect(list.IsList[int](studentList)).To(BeFalse())
			Expect(list.IsList[any](studentList)).To(BeFalse())
		})

		It("Should get elements", func() {
			Expect(studentList.GetAt(0).Name).To(Equal("Alice"))
			Expect(studentList.GetAt(1).Name).To(Equal("Bob"))
			Expect(studentList.GetAt(2).Name).To(Equal("Charlie"))
		})

		It("Should set elements", func() {
			studentList.SetAt(0, &Student{
				Name: "David",
			})
			studentList.SetAt(1, &Student{
				Name: "Eve",
			})
			studentList.SetAt(2, &Student{
				Name: "Frank",
			})

			Expect(studentList.GetAt(0).Name).To(Equal("David"))
			Expect(studentList.GetAt(1).Name).To(Equal("Eve"))
			Expect(studentList.GetAt(2).Name).To(Equal("Frank"))
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

				Expect(studentList.GetAt(index).Name).To(Equal(student.Name))
			})

			Expect(sum).To(Equal(3))
		})

		It("Should update elements", func() {
			var student = studentList.GetAt(0)
			student.SetName("Kafka")

			Expect(studentList.GetAt(0).Name).To(Equal("Kafka"))
		})

		It("Should try get in range", func() {
			var student, ok = studentList.TryGetAt(0)
			Expect(ok).To(BeTrue())
			Expect(student.Name).To(Equal("Alice"))

			student, ok = studentList.TryGetAt(1)
			Expect(ok).To(BeTrue())
			Expect(student.Name).To(Equal("Bob"))

			student, ok = studentList.TryGetAt(2)
			Expect(ok).To(BeTrue())
			Expect(student.Name).To(Equal("Charlie"))
		})

		It("Should try get out of range", func() {
			var student, ok = studentList.TryGetAt(-1)
			Expect(ok).To(BeFalse())
			Expect(student).To(Equal(utils.DefaultValue[*Student]()))

			student, ok = studentList.TryGetAt(3)
			Expect(ok).To(BeFalse())
			Expect(student).To(BeNil())
		})

		It("Should try set in range", func() {
			var ok = studentList.TrySetAt(0, &Student{
				Name: "David",
			})
			Expect(ok).To(BeTrue())
			Expect(studentList.GetAt(0).Name).To(Equal("David"))

			ok = studentList.TrySetAt(1, &Student{
				Name: "Eve",
			})
			Expect(ok).To(BeTrue())
			Expect(studentList.GetAt(1).Name).To(Equal("Eve"))

			ok = studentList.TrySetAt(2, &Student{
				Name: "Frank",
			})
			Expect(ok).To(BeTrue())
			Expect(studentList.GetAt(2).Name).To(Equal("Frank"))
		})

		It("Should try set out of range", func() {
			var ok = studentList.TrySetAt(-1, &Student{
				Name: "David",
			})
			Expect(ok).To(BeFalse())

			ok = studentList.TrySetAt(3, &Student{
				Name: "Eve",
			})
			Expect(ok).To(BeFalse())
		})

		It("Should try remove in range", func() {
			var student, ok = studentList.TryRemoveAt(0)
			Expect(ok).To(BeTrue())
			Expect(student.Name).To(Equal("Alice"))

			student, ok = studentList.TryRemoveAt(0)
			Expect(ok).To(BeTrue())
			Expect(student.Name).To(Equal("Bob"))

			student, ok = studentList.TryRemoveAt(0)
			Expect(ok).To(BeTrue())
			Expect(student.Name).To(Equal("Charlie"))

			Expect(studentList.Count()).To(Equal(0))
		})

		It("Should try remove out of range", func() {
			var student, ok = studentList.TryRemoveAt(-1)
			Expect(ok).To(BeFalse())
			Expect(student).To(Equal(utils.DefaultValue[*Student]()))

			student, ok = studentList.TryRemoveAt(11)
			Expect(ok).To(BeFalse())
			Expect(student).To(BeNil())
		})
	})
})

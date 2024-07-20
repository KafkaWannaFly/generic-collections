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

		It("Should remove elements", func() {
			var book0 = bookList.Remove(0)
			Expect(book0.Title).To(Equal("The Alchemist"))

			var book1 = bookList.Remove(0)
			Expect(book1.Title).To(Equal("The Little Prince"))

			var book2 = bookList.Remove(0)
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

				Expect(bookList.Get(index).Title).To(Equal(book.Title))
			})

			Expect(sum).To(Equal(3))
		})

		It("Should find elements", func() {
			var index = bookList.Find(func(book Book) bool {
				return book.Title == "The Little Prince"
			})

			Expect(index).To(Equal(1))
			Expect(bookList.Get(index).Title).To(Equal("The Little Prince"))
		})

		It("Should not find elements", func() {
			var index = bookList.Find(func(book Book) bool {
				return book.Title == "The Great Gatsby"
			})

			Expect(index).To(Equal(-1))
		})

		It("Should map to another type", func() {
			var titleList = bookList.Map(func(index int, book Book) any {
				return book.Title
			})

			Expect(titleList.Get(0)).To(Equal("The Alchemist"))
			Expect(titleList.Get(1)).To(Equal("The Little Prince"))
			Expect(titleList.Get(2)).To(Equal("The Catcher in the Rye"))

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
			Expect(group1943.Get(0).Title).To(Equal("The Little Prince"))

			var group1951 = groups.Get(1951)
			Expect(group1951.Count()).To(Equal(1))
			Expect(group1951.Get(0).Title).To(Equal("The Catcher in the Rye"))

			var group1988 = groups.Get(1988)
			Expect(group1988.Count()).To(Equal(1))
			Expect(group1988.Get(0).Title).To(Equal("The Alchemist"))
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

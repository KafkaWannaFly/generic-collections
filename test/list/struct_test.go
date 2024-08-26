package list

import (
	"github.com/KafkaWannaFly/generic-collections/list"
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
	})
})

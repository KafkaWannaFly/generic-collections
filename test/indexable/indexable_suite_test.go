package indexable_test

import (
	"github.com/KafkaWannaFly/generic-collections/interfaces"
	"github.com/KafkaWannaFly/generic-collections/linkedlist"
	"github.com/KafkaWannaFly/generic-collections/list"
	"github.com/KafkaWannaFly/generic-collections/queue"
	"github.com/KafkaWannaFly/generic-collections/stack"
	"github.com/KafkaWannaFly/generic-collections/utils"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// region Struct and Interface

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

// endregion

func TestIndexable(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Indexable Suite")
}

var _ = Describe("Test indexable collection", func() {

	When("Using integer", func() {
		integerLinkedList := linkedlist.From(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		Context("For LinkedList", testInteger(integerLinkedList))

		integerList := list.From(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		Context("For List", testInteger(integerList))

		integerStack := stack.From(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		Context("For Stack", testInteger(integerStack))

		integerQueue := queue.From(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		Context("For Queue", testInteger(integerQueue))
	})

	When("Using string", func() {
		stringLinkedList := linkedlist.From("a", "b", "c", "d", "e", "f", "g", "h", "i", "j")
		Context("For LinkedList", testString(stringLinkedList))

		stringList := list.From(stringLinkedList.ToSlice()...)
		Context("For List", testString(stringList))

		stringStack := stack.From(stringLinkedList.ToSlice()...)
		Context("For Stack", testString(stringStack))

		stringQueue := queue.From(stringLinkedList.ToSlice()...)
		Context("For Queue", testString(stringQueue))
	})

	When("Using struct", func() {
		bookList := list.New[Book]()
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

		Context("For List", testStruct(bookList))

		bookLinkedList := linkedlist.From(bookList.ToSlice()...)
		Context("For LinkedList", testStruct(bookLinkedList))

		bookStack := stack.From(bookList.ToSlice()...)
		Context("For Stack", testStruct(bookStack))

		bookQueue := queue.From(bookList.ToSlice()...)
		Context("For Queue", testStruct(bookQueue))
	})

	When("Using pointer", func() {
		studentList := list.New[*Student]()
		studentList.Add(
			&Student{
				Name: "Alice",
				Age:  20,
				GPA:  3.5,
			},
		).Add(
			&Student{
				Name: "Bob",
				Age:  21,
				GPA:  3.6,
			},
		).Add(
			&Student{
				Name: "Charlie",
				Age:  22,
				GPA:  3.7,
			},
		)

		Context("For List", testPointer(studentList))

		studentLinkedList := linkedlist.From(studentList.ToSlice()...)
		Context("For LinkedList", testPointer(studentLinkedList))

		studentStack := stack.From(studentList.ToSlice()...)
		Context("For Stack", testPointer(studentStack))

		studentQueue := queue.From(studentList.ToSlice()...)
		Context("For Queue", testPointer(studentQueue))
	})
})

func testInteger(collection interfaces.ICollection[int]) func() {
	return func() {
		var integerCollection interfaces.IIndexableCollection[int, int]

		BeforeEach(func() {
			integerCollection = collection.Clone().(any).(interfaces.IIndexableCollection[int, int])

			Expect(integerCollection.Count()).To(Equal(10))
			Expect(collection.ToSlice()).To(Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
		})

		It("Should be able to slice with internal result", func() {
			actualResult := integerCollection.Slice(0, 5)
			expectedResult := []int{1, 2, 3, 4, 5}

			Expect(actualResult.ToSlice()).To(Equal(expectedResult))
			Expect(actualResult.Count()).To(Equal(len(expectedResult)))
		})

		It("Should be able to slice with has wrap from begin", func() {
			actualResult := integerCollection.Slice(8, 5)
			expectedResult := []int{9, 10, 1, 2, 3}

			Expect(actualResult.ToSlice()).To(Equal(expectedResult))
			Expect(actualResult.Count()).To(Equal(len(expectedResult)))
		})

		It("Should be able to slice with panic when index out of range", func() {
			Expect(func() { integerCollection.Slice(999999999999999, 5) }).To(Panic())
			Expect(func() { integerCollection.Slice(-999999999999999, 5) }).To(Panic())
		})

		It("Should be able to slice with panic when length out of range", func() {
			Expect(func() { integerCollection.Slice(0, 999999999999999) }).To(Panic())
			Expect(func() { integerCollection.Slice(0, -999999999999999) }).To(Panic())
		})

		It("Should be able to get element by index", func() {
			for i := 0; i < integerCollection.Count(); i++ {
				Expect(integerCollection.GetAt(i)).To(Equal(i + 1))
			}
		})

		It("Should panic when get element out of range", func() {
			Expect(func() { integerCollection.GetAt(999999999999999) }).To(Panic())
		})

		It("Should try get element in range", func() {
			val, ok := integerCollection.TryGetAt(0)
			Expect(val).To(Equal(1))
			Expect(ok).To(BeTrue())
		})

		It("Should try get element out of range", func() {
			val, ok := integerCollection.TryGetAt(-99999999999999)
			Expect(val).To(Equal(utils.DefaultValue[int]()))
			Expect(ok).To(BeFalse())
		})

		It("Should be able to set element by index", func() {
			integerCollection.SetAt(0, 100)
			Expect(integerCollection.GetAt(0)).To(Equal(100))
		})

		It("Should panic when set element out of range", func() {
			Expect(func() { integerCollection.SetAt(999999999999999, 100) }).To(Panic())
		})

		It("Should try set element in range", func() {
			ok := integerCollection.TrySetAt(0, 100)
			Expect(ok).To(BeTrue())
			Expect(integerCollection.GetAt(0)).To(Equal(100))
		})

		It("Should try set element out of range", func() {
			ok := integerCollection.TrySetAt(-99999999999999, 100)
			Expect(ok).To(BeFalse())
		})

		It("Should add element to the beginning", func() {
			integerCollection.AddFirst(999)

			Expect(integerCollection.Count()).To(Equal(11))
			Expect(integerCollection.GetAt(0)).To(Equal(999))
		})

		It("Should add element to the end", func() {
			integerCollection.AddLast(999)

			Expect(integerCollection.Count()).To(Equal(11))
			Expect(integerCollection.GetAt(integerCollection.Count() - 1)).To(Equal(999))
		})

		It("Should add element before a certain index", func() {
			integerCollection.AddBefore(5, 999)

			Expect(integerCollection.Count()).To(Equal(11))
			Expect(integerCollection.GetAt(5)).To(Equal(999))

			integerCollection.AddBefore(0, 888)

			Expect(integerCollection.Count()).To(Equal(12))
			Expect(integerCollection.GetAt(0)).To(Equal(888))

			integerCollection.AddBefore(integerCollection.Count(), 777)

			Expect(integerCollection.Count()).To(Equal(13))
			Expect(integerCollection.GetAt(integerCollection.Count() - 1)).To(Equal(777))
		})

		It("Should be panic when add element before a certain index but out of range", func() {
			Expect(func() { integerCollection.AddBefore(999999999999999, 999) }).To(Panic())

			Expect(func() { integerCollection.AddBefore(-999999999999999, 999) }).To(Panic())
		})

		It("Should try to add element before a certain index", func() {
			ok := integerCollection.TryAddBefore(5, 999)
			Expect(ok).To(BeTrue())
			Expect(integerCollection.GetAt(5)).To(Equal(999))

			ok = integerCollection.TryAddBefore(0, 888)
			Expect(ok).To(BeTrue())
			Expect(integerCollection.GetAt(0)).To(Equal(888))

			ok = integerCollection.TryAddBefore(integerCollection.Count(), 777)
			Expect(ok).To(BeTrue())
			Expect(integerCollection.GetAt(integerCollection.Count() - 1)).To(Equal(777))
		})

		It("Should try to add element before a certain index but out of range", func() {
			ok := integerCollection.TryAddBefore(999999999999999, 999)
			Expect(ok).To(BeFalse())

			ok = integerCollection.TryAddBefore(-999999999999999, 999)
			Expect(ok).To(BeFalse())
		})

		It("Should add element after a certain index", func() {
			integerCollection.AddAfter(5, 999)

			Expect(integerCollection.Count()).To(Equal(11))
			Expect(integerCollection.GetAt(6)).To(Equal(999))

			integerCollection.AddAfter(0, 888)

			Expect(integerCollection.Count()).To(Equal(12))
			Expect(integerCollection.GetAt(1)).To(Equal(888))

			integerCollection.AddAfter(integerCollection.Count()-1, 777)

			Expect(integerCollection.Count()).To(Equal(13))
			Expect(integerCollection.GetAt(integerCollection.Count() - 1)).To(Equal(777))

			integerCollection.AddAfter(-1, 666)

			Expect(integerCollection.Count()).To(Equal(14))
			Expect(integerCollection.GetAt(0)).To(Equal(666))
		})

		It("Should try add element after a certain index", func() {
			ok := integerCollection.TryAddAfter(5, 999)
			Expect(ok).To(BeTrue())
			Expect(integerCollection.GetAt(6)).To(Equal(999))

			ok = integerCollection.TryAddAfter(0, 888)
			Expect(ok).To(BeTrue())
			Expect(integerCollection.GetAt(1)).To(Equal(888))

			ok = integerCollection.TryAddAfter(integerCollection.Count()-1, 777)
			Expect(ok).To(BeTrue())
			Expect(integerCollection.GetAt(integerCollection.Count() - 1)).To(Equal(777))

			ok = integerCollection.TryAddAfter(-1, 666)
			Expect(ok).To(BeTrue())
			Expect(integerCollection.GetAt(0)).To(Equal(666))
		})

		It("Should try add element after a certain index but out of range", func() {
			ok := integerCollection.TryAddAfter(999999999999999, 999)
			Expect(ok).To(BeFalse())

			ok = integerCollection.TryAddAfter(-999999999999999, 999)
			Expect(ok).To(BeFalse())
		})

		It("Should remove first element", func() {
			first := integerCollection.RemoveFirst()
			Expect(integerCollection.Count()).To(Equal(9))
			Expect(first).To(Equal(1))
			Expect(integerCollection.GetAt(0)).To(Equal(2))

			first = integerCollection.RemoveFirst()
			Expect(integerCollection.Count()).To(Equal(8))
			Expect(first).To(Equal(2))
			Expect(integerCollection.GetAt(0)).To(Equal(3))

			first = integerCollection.RemoveFirst()
			Expect(integerCollection.Count()).To(Equal(7))
			Expect(first).To(Equal(3))
			Expect(integerCollection.GetAt(0)).To(Equal(4))
		})

		It("Should remove last element", func() {
			last := integerCollection.RemoveLast()
			Expect(integerCollection.Count()).To(Equal(9))
			Expect(last).To(Equal(10))
			Expect(integerCollection.GetAt(integerCollection.Count() - 1)).To(Equal(9))

			last = integerCollection.RemoveLast()
			Expect(integerCollection.Count()).To(Equal(8))
			Expect(last).To(Equal(9))
			Expect(integerCollection.GetAt(integerCollection.Count() - 1)).To(Equal(8))

			last = integerCollection.RemoveLast()
			Expect(integerCollection.Count()).To(Equal(7))
			Expect(last).To(Equal(8))
			Expect(integerCollection.GetAt(integerCollection.Count() - 1)).To(Equal(7))
		})

		It("Should remove element by index", func() {
			integerCollection.RemoveAt(0)
			Expect(integerCollection.Count()).To(Equal(9))
			Expect(integerCollection.GetAt(0)).To(Equal(2))

			integerCollection.RemoveAt(integerCollection.Count() - 1)
			Expect(integerCollection.Count()).To(Equal(8))
			Expect(integerCollection.GetAt(integerCollection.Count() - 1)).To(Equal(9))
		})

		It("Should panic when remove element out of range", func() {
			Expect(func() { integerCollection.RemoveAt(999999999999999) }).To(Panic())

			Expect(func() { integerCollection.RemoveAt(-999999999999999) }).To(Panic())
		})

		It("Should try remove element by index", func() {
			val, ok := integerCollection.TryRemoveAt(0)
			Expect(integerCollection.Count()).To(Equal(9))
			Expect(val).To(Equal(1))
			Expect(ok).To(BeTrue())
		})

		It("Should try remove element by index but out of range", func() {
			val, ok := integerCollection.TryRemoveAt(999999999999999)
			Expect(val).To(Equal(utils.DefaultValue[int]()))
			Expect(ok).To(BeFalse())

			val, ok = integerCollection.TryRemoveAt(-999999999999999)
			Expect(val).To(Equal(utils.DefaultValue[int]()))
			Expect(ok).To(BeFalse())
		})

		It("Should find the first element based on predicate", func() {
			index := integerCollection.FindFirst(func(i int, val int) bool { return val == 5 })
			Expect(index).To(Equal(4))

			index = integerCollection.FindFirst(func(i int, val int) bool { return val == 999 })
			Expect(index).To(Equal(-1))
		})

		It("Should find the last element based on predicate", func() {
			index := integerCollection.FindLast(func(i int, val int) bool { return val == 5 })
			Expect(index).To(Equal(4))

			index = integerCollection.FindLast(func(i int, val int) bool { return val == 999 })
			Expect(index).To(Equal(-1))
		})

		It("Should find all elements based on predicate", func() {
			indexes := integerCollection.FindAll(func(i int, val int) bool { return val > 5 })
			Expect(indexes).To(Equal([]int{5, 6, 7, 8, 9}))

			indexes = integerCollection.FindAll(func(i int, val int) bool { return val < 0 })
			Expect(indexes).To(Equal([]int{}))
		})
	}
}

func testString(collection interfaces.ICollection[string]) func() {
	return func() {
		var stringCollection interfaces.IIndexableCollection[int, string]

		BeforeEach(func() {
			stringCollection = collection.Clone().(any).(interfaces.IIndexableCollection[int, string])

			Expect(stringCollection.Count()).To(Equal(10))
		})

		It("Should add an element at the head", func() {
			stringCollection.AddFirst("z")

			Expect(stringCollection.Count()).To(Equal(11))
			Expect(stringCollection.GetAt(0)).To(Equal("z"))
		})

		It("Should add an element at the tail", func() {
			stringCollection.AddLast("z")

			Expect(stringCollection.Count()).To(Equal(11))
			Expect(stringCollection.GetAt(10)).To(Equal("z"))
		})

		It("Should add before an element", func() {
			stringCollection.AddBefore(0, "head")
			stringCollection.AddBefore(stringCollection.Count(), "tail")

			Expect(stringCollection.Count()).To(Equal(12))
			Expect(stringCollection.GetAt(0)).To(Equal("head"))
			Expect(stringCollection.GetAt(stringCollection.Count() - 1)).To(Equal("tail"))

			stringCollection.AddBefore(5, "middle")

			Expect(stringCollection.Count()).To(Equal(13))
			Expect(stringCollection.GetAt(5)).To(Equal("middle"))
		})

		It("Should add after an element", func() {
			stringCollection.AddAfter(-1, "head")
			stringCollection.AddAfter(stringCollection.Count()-1, "tail")

			Expect(stringCollection.Count()).To(Equal(12))
			Expect(stringCollection.GetAt(0)).To(Equal("head"))
			Expect(stringCollection.GetAt(stringCollection.Count() - 1)).To(Equal("tail"))

			stringCollection.AddAfter(5, "middle")

			Expect(stringCollection.Count()).To(Equal(13))
			Expect(stringCollection.GetAt(6)).To(Equal("middle"))
		})

		It("Should try add before an element in range", func() {
			ok := stringCollection.TryAddBefore(0, "head")
			Expect(ok).To(BeTrue())
			Expect(stringCollection.Count()).To(Equal(11))
			Expect(stringCollection.GetAt(0)).To(Equal("head"))
		})

		It("Should try add before an element out of range", func() {
			ok := stringCollection.TryAddBefore(-1, "head")
			Expect(ok).To(BeFalse())
			Expect(stringCollection.Count()).To(Equal(10))
		})

		It("Should try add after an element in range", func() {
			ok := stringCollection.TryAddAfter(-1, "head")
			Expect(ok).To(BeTrue())
			Expect(stringCollection.Count()).To(Equal(11))
			Expect(stringCollection.GetAt(0)).To(Equal("head"))
		})

		It("Should try add after an element out of range", func() {
			ok := stringCollection.TryAddAfter(stringCollection.Count(), "tail")
			Expect(ok).To(BeFalse())
			Expect(stringCollection.Count()).To(Equal(10))
		})

		It("Should remove first item", func() {
			stringCollection.RemoveFirst()

			Expect(stringCollection.Count()).To(Equal(9))
			Expect(stringCollection.GetAt(0)).To(Equal("b"))
		})

		It("Should remove last item", func() {
			stringCollection.RemoveLast()

			Expect(stringCollection.Count()).To(Equal(9))
			Expect(stringCollection.GetAt(8)).To(Equal("i"))
		})

		It("Should get an element", func() {
			Expect(stringCollection.GetAt(0)).To(Equal("a"))
			Expect(stringCollection.GetAt(1)).To(Equal("b"))
			Expect(stringCollection.GetAt(2)).To(Equal("c"))
			Expect(stringCollection.GetAt(3)).To(Equal("d"))
			Expect(stringCollection.GetAt(4)).To(Equal("e"))
			Expect(stringCollection.GetAt(5)).To(Equal("f"))
			Expect(stringCollection.GetAt(6)).To(Equal("g"))
			Expect(stringCollection.GetAt(7)).To(Equal("h"))
			Expect(stringCollection.GetAt(8)).To(Equal("i"))
			Expect(stringCollection.GetAt(9)).To(Equal("j"))
		})

		It("Should set an element", func() {
			stringCollection.SetAt(0, "aa")
			stringCollection.SetAt(1, "bb")
			stringCollection.SetAt(2, "cc")
			stringCollection.SetAt(3, "dd")
			stringCollection.SetAt(4, "ee")
			stringCollection.SetAt(5, "ff")
			stringCollection.SetAt(6, "gg")
			stringCollection.SetAt(7, "hh")
			stringCollection.SetAt(8, "ii")
			stringCollection.SetAt(9, "jj")

			Expect(stringCollection.GetAt(0)).To(Equal("aa"))
			Expect(stringCollection.GetAt(1)).To(Equal("bb"))
			Expect(stringCollection.GetAt(2)).To(Equal("cc"))
			Expect(stringCollection.GetAt(3)).To(Equal("dd"))
			Expect(stringCollection.GetAt(4)).To(Equal("ee"))
			Expect(stringCollection.GetAt(5)).To(Equal("ff"))
			Expect(stringCollection.GetAt(6)).To(Equal("gg"))
			Expect(stringCollection.GetAt(7)).To(Equal("hh"))
			Expect(stringCollection.GetAt(8)).To(Equal("ii"))
			Expect(stringCollection.GetAt(9)).To(Equal("jj"))
		})

		It("Should remove an element", func() {
			first := stringCollection.GetAt(0)
			firstRemoveAt := stringCollection.RemoveAt(0)

			Expect(firstRemoveAt).To(Equal(first))

			lastIndex := stringCollection.Count() - 1
			last := stringCollection.GetAt(lastIndex)
			lastRemoveAt := stringCollection.RemoveAt(lastIndex)

			Expect(lastRemoveAt).To(Equal(last))

			Expect(stringCollection.Count()).To(Equal(8))

			Expect(stringCollection.GetAt(0)).To(Equal("b"))
			Expect(stringCollection.GetAt(1)).To(Equal("c"))
			Expect(stringCollection.GetAt(2)).To(Equal("d"))
			Expect(stringCollection.GetAt(3)).To(Equal("e"))
			Expect(stringCollection.GetAt(4)).To(Equal("f"))
			Expect(stringCollection.GetAt(5)).To(Equal("g"))
			Expect(stringCollection.GetAt(6)).To(Equal("h"))
			Expect(stringCollection.GetAt(7)).To(Equal("i"))
		})

		It("Should find an element", func() {
			var index = stringCollection.FindFirst(func(_ int, element string) bool {
				return element == "e"
			})

			Expect(index).To(Equal(4))
			Expect(stringCollection.GetAt(index)).To(Equal("e"))

			firstIndex := stringCollection.FindFirst(func(_ int, element string) bool {
				return element == "b"
			})

			Expect(firstIndex).To(Equal(1))
			Expect(stringCollection.GetAt(firstIndex)).To(Equal("b"))

			lastIndex := stringCollection.FindLast(func(_ int, element string) bool {
				return element >= "a"
			})
			Expect(lastIndex).To(Equal(9))
			Expect(stringCollection.GetAt(lastIndex)).To(Equal("j"))
		})

		It("Should not find an element", func() {
			var index = stringCollection.FindFirst(func(_ int, element string) bool {
				return element == "z"
			})

			Expect(index).To(Equal(-1))

			firstIndex := stringCollection.FindFirst(func(_ int, element string) bool {
				return element == "z"
			})

			Expect(firstIndex).To(Equal(-1))

			lastIndex := stringCollection.FindLast(func(_ int, element string) bool {
				return element == "z"
			})

			Expect(lastIndex).To(Equal(-1))
		})

		It("Should try get element in range", func() {
			val, ok := stringCollection.TryGetAt(0)
			Expect(val).To(Equal("a"))
			Expect(ok).To(BeTrue())

			val, ok = stringCollection.TryGetAt(1)
			Expect(val).To(Equal("b"))
			Expect(ok).To(BeTrue())

			val, ok = stringCollection.TryGetAt(9)
			Expect(val).To(Equal("j"))
			Expect(ok).To(BeTrue())
		})

		It("Should try get element out of range", func() {
			val, ok := stringCollection.TryGetAt(-1)
			Expect(val).To(Equal(""))
			Expect(ok).To(BeFalse())

			val, ok = stringCollection.TryGetAt(10)
			Expect(val).To(Equal(""))
			Expect(ok).To(BeFalse())
		})

		It("Should try set element in range", func() {
			ok := stringCollection.TrySetAt(0, "aa")
			Expect(ok).To(BeTrue())
			Expect(stringCollection.GetAt(0)).To(Equal("aa"))

			ok = stringCollection.TrySetAt(1, "bb")
			Expect(ok).To(BeTrue())
			Expect(stringCollection.GetAt(1)).To(Equal("bb"))

			ok = stringCollection.TrySetAt(9, "jj")
			Expect(ok).To(BeTrue())
			Expect(stringCollection.GetAt(9)).To(Equal("jj"))
		})

		It("Should try set element out of range", func() {
			ok := stringCollection.TrySetAt(-1, "0")
			Expect(ok).To(BeFalse())

			ok = stringCollection.TrySetAt(10, "k")
			Expect(ok).To(BeFalse())
		})

		It("Should try remove element in range", func() {
			val, ok := stringCollection.TryRemoveAt(0)
			Expect(val).To(Equal("a"))
			Expect(ok).To(BeTrue())
			Expect(stringCollection.Count()).To(Equal(9))

			val, ok = stringCollection.TryRemoveAt(8)
			Expect(val).To(Equal("j"))
			Expect(ok).To(BeTrue())
			Expect(stringCollection.Count()).To(Equal(8))
		})

		It("Should try remove element out of range", func() {
			val, ok := stringCollection.TryRemoveAt(-1)
			Expect(val).To(Equal(""))
			Expect(ok).To(BeFalse())
			Expect(stringCollection.Count()).To(Equal(10))

			val, ok = stringCollection.TryRemoveAt(10)
			Expect(val).To(Equal(""))
			Expect(ok).To(BeFalse())
			Expect(stringCollection.Count()).To(Equal(10))
		})
	}
}

func testStruct(collection interfaces.ICollection[Book]) func() {
	return func() {
		var bookCollection interfaces.IIndexableCollection[int, Book]

		BeforeEach(func() {
			bookCollection = collection.Clone().(any).(interfaces.IIndexableCollection[int, Book])

			Expect(bookCollection.Count()).To(Equal(3))
		})

		It("Should get elements", func() {
			Expect(bookCollection.GetAt(0).Title).To(Equal("The Alchemist"))
			Expect(bookCollection.GetAt(1).Title).To(Equal("The Little Prince"))
			Expect(bookCollection.GetAt(2).Title).To(Equal("The Catcher in the Rye"))
		})

		It("Should set elements", func() {
			bookCollection.SetAt(0, Book{
				Title: "The Great Gatsby",
			})
			bookCollection.SetAt(1, Book{
				Title: "To Kill a Mockingbird",
			})
			bookCollection.SetAt(2, Book{
				Title: "1984",
			})

			Expect(bookCollection.GetAt(0).Title).To(Equal("The Great Gatsby"))
			Expect(bookCollection.GetAt(1).Title).To(Equal("To Kill a Mockingbird"))
			Expect(bookCollection.GetAt(2).Title).To(Equal("1984"))
		})

		It("Should remove elements", func() {
			var book0 = bookCollection.RemoveAt(0)
			Expect(book0.Title).To(Equal("The Alchemist"))

			var book1 = bookCollection.RemoveAt(0)
			Expect(book1.Title).To(Equal("The Little Prince"))

			var book2 = bookCollection.RemoveAt(0)
			Expect(book2.Title).To(Equal("The Catcher in the Rye"))

			Expect(bookCollection.Count()).To(Equal(0))
		})

		It("Should convert to slice", func() {
			var bookSlice = bookCollection.ToSlice()
			Expect(bookSlice[0].Title).To(Equal("The Alchemist"))
			Expect(bookSlice[1].Title).To(Equal("The Little Prince"))
			Expect(bookSlice[2].Title).To(Equal("The Catcher in the Rye"))

			Expect(len(bookSlice)).To(Equal(3))
		})

		It("Should iterate over elements", func() {
			var sum int
			bookCollection.ForEach(func(index int, book Book) {
				sum += 1

				Expect(bookCollection.GetAt(index).Title).To(Equal(book.Title))
			})

			Expect(sum).To(Equal(3))
		})

		It("Should find elements", func() {
			var index = bookCollection.FindFirst(func(_ int, book Book) bool {
				return book.Title == "The Little Prince"
			})

			Expect(index).To(Equal(1))
			Expect(bookCollection.GetAt(index).Title).To(Equal("The Little Prince"))
		})

		It("Should not find elements", func() {
			var index = bookCollection.FindFirst(func(_ int, book Book) bool {
				return book.Title == "The Great Gatsby"
			})

			Expect(index).To(Equal(-1))
		})

		It("Should try get elements in range", func() {
			var book, ok = bookCollection.TryGetAt(0)
			Expect(ok).To(BeTrue())
			Expect(book.Title).To(Equal("The Alchemist"))

			book, ok = bookCollection.TryGetAt(1)
			Expect(ok).To(BeTrue())
			Expect(book.Title).To(Equal("The Little Prince"))

			book, ok = bookCollection.TryGetAt(2)
			Expect(ok).To(BeTrue())
			Expect(book.Title).To(Equal("The Catcher in the Rye"))
		})

		It("Should try get elements out of range", func() {
			var book, ok = bookCollection.TryGetAt(-1)
			Expect(ok).To(BeFalse())
			Expect(book).To(Equal(utils.DefaultValue[Book]()))

			book, ok = bookCollection.TryGetAt(3)
			Expect(ok).To(BeFalse())
			var defaultBook Book
			Expect(book).To(Equal(defaultBook))
		})

		It("Should try set elements in range", func() {
			var ok = bookCollection.TrySetAt(0, Book{
				Title: "The Great Gatsby",
			})
			Expect(ok).To(BeTrue())
			Expect(bookCollection.GetAt(0).Title).To(Equal("The Great Gatsby"))

			ok = bookCollection.TrySetAt(1, Book{
				Title: "To Kill a Mockingbird",
			})
			Expect(ok).To(BeTrue())
			Expect(bookCollection.GetAt(1).Title).To(Equal("To Kill a Mockingbird"))

			ok = bookCollection.TrySetAt(2, Book{
				Title: "1984",
			})
			Expect(ok).To(BeTrue())
			Expect(bookCollection.GetAt(2).Title).To(Equal("1984"))
		})

		It("Should try set elements out of range", func() {
			var ok = bookCollection.TrySetAt(-1, Book{
				Title: "The Great Gatsby",
			})
			Expect(ok).To(BeFalse())

			ok = bookCollection.TrySetAt(3, Book{
				Title: "To Kill a Mockingbird",
			})
			Expect(ok).To(BeFalse())
		})

		It("Should try remove elements in range", func() {
			var book, ok = bookCollection.TryRemoveAt(0)
			Expect(ok).To(BeTrue())
			Expect(book.Title).To(Equal("The Alchemist"))

			book, ok = bookCollection.TryRemoveAt(0)
			Expect(ok).To(BeTrue())
			Expect(book.Title).To(Equal("The Little Prince"))

			book, ok = bookCollection.TryRemoveAt(0)
			Expect(ok).To(BeTrue())
			Expect(book.Title).To(Equal("The Catcher in the Rye"))

			Expect(bookCollection.Count()).To(Equal(0))
		})

		It("Should try remove elements out of range", func() {
			var book, ok = bookCollection.TryRemoveAt(-1)
			Expect(ok).To(BeFalse())
			Expect(book).To(Equal(utils.DefaultValue[Book]()))

			book, ok = bookCollection.TryRemoveAt(11)
			Expect(ok).To(BeFalse())
			Expect(book).To(Equal(utils.DefaultValue[Book]()))
		})
	}
}

func testPointer(collection interfaces.ICollection[*Student]) func() {
	return func() {
		var studentCollection interfaces.IIndexableCollection[int, *Student]

		BeforeEach(func() {
			studentCollection = collection.Clone().(any).(interfaces.IIndexableCollection[int, *Student])

			Expect(studentCollection.Count()).To(Equal(3))
		})

		It("Should get elements", func() {
			Expect(studentCollection.GetAt(0).Name).To(Equal("Alice"))
			Expect(studentCollection.GetAt(1).Name).To(Equal("Bob"))
			Expect(studentCollection.GetAt(2).Name).To(Equal("Charlie"))
		})

		It("Should set elements", func() {
			studentCollection.SetAt(0, &Student{
				Name: "David",
			})
			studentCollection.SetAt(1, &Student{
				Name: "Eve",
			})
			studentCollection.SetAt(2, &Student{
				Name: "Frank",
			})

			Expect(studentCollection.GetAt(0).Name).To(Equal("David"))
			Expect(studentCollection.GetAt(1).Name).To(Equal("Eve"))
			Expect(studentCollection.GetAt(2).Name).To(Equal("Frank"))
		})

		It("Should convert to slice", func() {
			var studentSlice = studentCollection.ToSlice()
			Expect(studentSlice[0].Name).To(Equal("Alice"))
			Expect(studentSlice[1].Name).To(Equal("Bob"))
			Expect(studentSlice[2].Name).To(Equal("Charlie"))

			Expect(len(studentSlice)).To(Equal(3))
		})

		It("Should iterate over elements", func() {
			var sum int
			studentCollection.ForEach(func(index int, student *Student) {
				sum += 1

				Expect(studentCollection.GetAt(index).Name).To(Equal(student.Name))
			})

			Expect(sum).To(Equal(3))
		})

		It("Should try get in range", func() {
			var student, ok = studentCollection.TryGetAt(0)
			Expect(ok).To(BeTrue())
			Expect(student.Name).To(Equal("Alice"))

			student, ok = studentCollection.TryGetAt(1)
			Expect(ok).To(BeTrue())
			Expect(student.Name).To(Equal("Bob"))

			student, ok = studentCollection.TryGetAt(2)
			Expect(ok).To(BeTrue())
			Expect(student.Name).To(Equal("Charlie"))
		})

		It("Should try get out of range", func() {
			var student, ok = studentCollection.TryGetAt(-1)
			Expect(ok).To(BeFalse())
			Expect(student).To(Equal(utils.DefaultValue[*Student]()))

			student, ok = studentCollection.TryGetAt(3)
			Expect(ok).To(BeFalse())
			Expect(student).To(BeNil())
		})

		It("Should try set in range", func() {
			var ok = studentCollection.TrySetAt(0, &Student{
				Name: "David",
			})
			Expect(ok).To(BeTrue())
			Expect(studentCollection.GetAt(0).Name).To(Equal("David"))

			ok = studentCollection.TrySetAt(1, &Student{
				Name: "Eve",
			})
			Expect(ok).To(BeTrue())
			Expect(studentCollection.GetAt(1).Name).To(Equal("Eve"))

			ok = studentCollection.TrySetAt(2, &Student{
				Name: "Frank",
			})
			Expect(ok).To(BeTrue())
			Expect(studentCollection.GetAt(2).Name).To(Equal("Frank"))
		})

		It("Should try set out of range", func() {
			var ok = studentCollection.TrySetAt(-1, &Student{
				Name: "David",
			})
			Expect(ok).To(BeFalse())

			ok = studentCollection.TrySetAt(3, &Student{
				Name: "Eve",
			})
			Expect(ok).To(BeFalse())
		})

		It("Should try remove in range", func() {
			var student, ok = studentCollection.TryRemoveAt(0)
			Expect(ok).To(BeTrue())
			Expect(student.Name).To(Equal("Alice"))

			student, ok = studentCollection.TryRemoveAt(0)
			Expect(ok).To(BeTrue())
			Expect(student.Name).To(Equal("Bob"))

			student, ok = studentCollection.TryRemoveAt(0)
			Expect(ok).To(BeTrue())
			Expect(student.Name).To(Equal("Charlie"))

			Expect(studentCollection.Count()).To(Equal(0))
		})

		It("Should try remove out of range", func() {
			var student, ok = studentCollection.TryRemoveAt(-1)
			Expect(ok).To(BeFalse())
			Expect(student).To(Equal(utils.DefaultValue[*Student]()))

			student, ok = studentCollection.TryRemoveAt(11)
			Expect(ok).To(BeFalse())
			Expect(student).To(BeNil())
		})
	}
}

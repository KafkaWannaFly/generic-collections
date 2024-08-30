# Set

`Set` is unordered and implements the [`ICollection`](../icollection.md) interface. It represents a collection of unique elements.

## Example

### Basic

The basic usage is very similar to [list](../list/ "mention") because they both implement [icollection.md](../icollection.md "mention")

{% code overflow="wrap" %}
```go
package main

import (
	"fmt"
	"github.com/KafkaWannaFly/generic-collections/list"
	"github.com/KafkaWannaFly/generic-collections/set"
)

func main() {
	integerSet := set.From(1, 2, 3, 4, 5)

	// Loop through the set
	// Set is unordered, so the index is always 0
	// The order of the elements is not guaranteed
	integerSet.ForEach(func(index int, item int) {
		fmt.Printf("Index: %d, Item: %d\n", index, item)
	})

	// Output: 1 2 3 4 5 6 7 8
	// 1 2 3 4 5 are already in the set
	integerSet.Add(1).Add(2).Add(3).Add(4).Add(5).Add(6).Add(7).Add(8)

	integerSet.AddAll(
		set.From(9, 10, 11, 12, 13, 14, 15),
	)

	fmt.Printf("Assert integer set %t\n", set.IsSet[int](integerSet))   // true
	fmt.Printf("Assert string set %t\n", set.IsSet[string](integerSet)) // false

	fmt.Printf("Contains 1: %t\n", integerSet.Has(1))     // true
	fmt.Printf("Contains 100: %t\n", integerSet.Has(100)) // false

	fmt.Printf(
		"Contains all 1, 2, 3, 4, 5: %t\n",
		integerSet.HasAll(list.From(1, 2, 3, 4, 5)),
	) // true
}
```
{% endcode %}

### Set Specific Methods

<figure><img src="../../.gitbook/assets/set functions (1).png" alt=""><figcaption></figcaption></figure>

```go
package main

import "github.com/KafkaWannaFly/generic-collections/set"

type Book struct {
	Title  string
	Author string
}

// Set will use this method to compare equality between books
// If you don't define this method, set will try to convert the struct to a string and compare them
func (receiver Book) HashCode() string {
	return receiver.Title + receiver.Author
}

func main() {
	norwegianWood := Book{Title: "Norwegian Wood", Author: "Haruki Murakami"}
	kafkaOnTheShore := Book{Title: "Kafka on the Shore", Author: "Haruki Murakami"}
	wildSheepChase := Book{Title: "A Wild Sheep Chase", Author: "Haruki Murakami"}

	kafkaOnTheShoreCopy := Book{Title: "Kafka on the Shore", Author: "Haruki Murakami"}
	chronicleOfWindUpBird := Book{Title: "Chronicle of Wind-Up Bird", Author: "Haruki Murakami"}

	bookSet1 := set.From(norwegianWood, kafkaOnTheShore, wildSheepChase)
	bookSet2 := set.From(kafkaOnTheShoreCopy, chronicleOfWindUpBird)

	// Union
	union := bookSet1.Union(bookSet2) // {Norwegian Wood, Kafka on the Shore, A Wild Sheep Chase, Chronicle of Wind-Up Bird}

	// Intersection
	intersection := bookSet1.Intersect(bookSet2) // {Kafka on the Shore}

	// Difference
	difference := bookSet1.Difference(bookSet2) // {Norwegian Wood, A Wild Sheep Chase}

	// Symmetric Difference
	symmetricDifference := bookSet1.SymmetricDifference(bookSet2) // {Norwegian Wood, A Wild Sheep Chase, Chronicle of Wind-Up Bird}
}

```

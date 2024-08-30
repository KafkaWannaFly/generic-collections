# List

`List` is a collection with index. It implements [`IIndexableCollection`](iindexablecollection.md).

## Example

### Create a List

```go
package main

import (
	"github.com/KafkaWannaFly/generic-collections/list"
)

func main() {
	// You can create a new list from a slice
	integerList := list.From(1, 2, 3, 4, 5)

	// You also can chaining the Add method
	chainedList := list.New[int]()
	chainedList.Add(1).Add(2).Add(3).Add(4).Add(5)

	// You can add all elements from chainedList to integerList
	integerList.AddAll(chainedList)
}

```

### Get and Set

```go
package main

import (
	"fmt"
	"github.com/KafkaWannaFly/generic-collections/list"
)

type Warrior struct {
	Name      string
	BodyCount float64
}

func main() {
	eren := Warrior{"Eren", 1000}
	mikasa := Warrior{"Mikasa", 2000}
	armin := Warrior{"Armin", 3000}
	levi := Warrior{"Levi", 4000}
	hange := Warrior{"Hange", 5000}

	warriors := list.From(eren, mikasa, armin, levi, hange)

	// Loop through the warriors and print their names and body counts.
	warriors.ForEach(func(index int, warrior Warrior) {
		println(warrior.Name, warrior.BodyCount)
	})

	// Get the first warrior.
	firstItem := warriors.GetAt(0)
	// First warrior: {Eren 1000}
	fmt.Printf("First warrior: %v\n", firstItem)

	// Safe get operation.
	threeHundreds, ok := warriors.TryGetAt(300)
	// 300th warrior: { 0} false
	fmt.Printf("300th warrior: %v %t\n", threeHundreds, ok)
	
	// Set the 1st warrior to a new warrior.
	warriors.SetAt(0, Warrior{"Jean", 6000})
	// Safe set operation.
	warriors.TrySetAt(300, Warrior{"Connie", 7000}) // This will not change anything.
}

```

### Map and Sum

```go
package main

import (
	"fmt"
	"github.com/KafkaWannaFly/generic-collections/list"
)

type Warrior struct {
	Name      string
	BodyCount float64
}

func main() {
	eren := Warrior{"Eren", 1000}
	mikasa := Warrior{"Mikasa", 2000}
	armin := Warrior{"Armin", 3000}
	levi := Warrior{"Levi", 4000}
	hange := Warrior{"Hange", 5000}

	warriors := list.From(eren, mikasa, armin, levi, hange)

	// Due to Go limitation with generic. Data type is *list.List[any]
	anyBodies := warriors.Map(func(index int, warrior Warrior) any {
		return warrior.BodyCount
	})

	// To have better type safety, you can use the following approach
	// Data type is *list.List[float64]
	bodyCount := list.Map(warriors, func(index int, warrior Warrior) float64 {
		return warrior.BodyCount
	})

	// Use can use package functions to perform operations
	totalBodies := list.Reduce(bodyCount, func(accumulator float64, income float64) float64 {
		return accumulator + income
	}, 0)

	// Or you can use list methods
	// But have to cast the result to the desired type
	totalBodies = anyBodies.Reduce(func(acc any, income any) any {
		return acc.(float64) + income.(float64)
	}, 0.0).(float64)

	// Check for data type
	fmt.Println(list.IsList[Warrior](warriors)) // true

	fmt.Println(list.IsList[any](anyBodies))     // true
	fmt.Println(list.IsList[float64](anyBodies)) // false

	fmt.Println(list.IsList[float64](bodyCount)) // true

	// Total body count: 15000.000000
	fmt.Printf("Total body count: %f\n", totalBodies)
}

```

### Grouping

```go
package main

import (
	"fmt"
	"github.com/KafkaWannaFly/generic-collections/list"
)

type Warrior struct {
	Name      string
	BodyCount float64
}

func main() {
	eren := Warrior{"Eren", 1000}
	mikasa := Warrior{"Mikasa", 2000}
	armin := Warrior{"Armin", 3000}
	levi := Warrior{"Levi", 4000}
	hange := Warrior{"Hange", 5000}

	warriors := list.From(eren, mikasa, armin, levi, hange)

	// Group by BodyCount. Low tier warriors have BodyCount < 3000. High tier warriors have BodyCount >= 3000
	// We will have 2 groups: Low and High
	// Output will be a *hashmap.HashMap[any, *list.List[Warrior]]
	// 2 groups means 2 keys: Low and High
	// Once again, due to Go limitations, key type is any
	anyTier := warriors.GroupBy(func(warrior Warrior) any {
		if warrior.BodyCount < 3000 {
			return "Low"
		} else {
			return "High"
		}
	})

	// Use package function to have better type assertion
	// Data-type *hashmap.HashMap[string, *list.List[Warrior]]
	tier := list.GroupBy(warriors, func(warrior Warrior) string {
		if warrior.BodyCount < 3000 {
			return "Low"
		} else {
			return "High"
		}
	})

	// Low tier warriors: &{[{Eren 1000} {Mikasa 2000}] 2}
	fmt.Printf("Low tier warriors: %v\n", tier.Get("Low"))
	
	// High tier warriors: &{[{Armin 3000} {Levi 4000} {Hange 5000}] 3}
	fmt.Printf("High tier warriors: %v\n", anyTier.Get("High"))
}

```

### Filtering

```go
package main

import (
	"fmt"
	"github.com/KafkaWannaFly/generic-collections/list"
)

type Warrior struct {
	Name      string
	BodyCount float64
}

func main() {
	eren := Warrior{"Eren", 1000}
	mikasa := Warrior{"Mikasa", 2000}
	armin := Warrior{"Armin", 3000}
	levi := Warrior{"Levi", 4000}
	hange := Warrior{"Hange", 5000}

	warriors := list.From(eren, mikasa, armin, levi, hange)

	// Filter warriors with body count greater than 2000
	// Return a new list. Original list is not modified
	goodWarriors := warriors.Filter(func(warrior Warrior) bool {
		return warrior.BodyCount > 2000
	})

	// Good warriors: &{[{Armin 3000} {Levi 4000} {Hange 5000}] 3}
	fmt.Printf("Good warriors: %v\n", goodWarriors)
}

```

### Slice a List

```go
package main

import (
	"fmt"
	"github.com/KafkaWannaFly/generic-collections/list"
)

type Warrior struct {
	Name      string
	BodyCount float64
}

func main() {
	eren := Warrior{"Eren", 1000}
	mikasa := Warrior{"Mikasa", 2000}
	armin := Warrior{"Armin", 3000}
	levi := Warrior{"Levi", 4000}
	hange := Warrior{"Hange", 5000}

	warriors := list.From(eren, mikasa, armin, levi, hange)

	// Cut the list to the first two elements.
	teamEren := warriors.Slice(0, 2)
	// Small team: &{[{Eren 1000} {Mikasa 2000}] 2}
	fmt.Printf("Small team: %v\n", teamEren)

	// Cut the list to the last two elements.
	teamLevi := warriors.Slice(3, 2)
	// Levi's team: &{[{Levi 4000} {Hange 5000}] 2}
	fmt.Printf("Levi's team: %v\n", teamLevi)

	// If length is surpassed the end, it will loop back to the beginning.
	merryGoRound := warriors.Slice(3, 5)
	// &{[{Levi 4000} {Hange 5000} {Eren 1000} {Mikasa 2000} {Armin 3000}] 5}
	fmt.Println(merryGoRound)
}

```

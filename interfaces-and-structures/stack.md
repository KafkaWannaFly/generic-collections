# Stack

`Stack` implements [iindexablecollection.md](iindexablecollection.md "mention") and provides specific methods for LIFO (Last In First Out) operations. It has a `LinkedList` under the hood, so `Stack` is a good fit for intensive `Push` and `Pop` operations.

<figure><img src="../.gitbook/assets/925177.jpg" alt=""><figcaption></figcaption></figure>

```go
package main

import (
	"fmt"
	"github.com/KafkaWannaFly/generic-collections/stack"
)

func main() {
	stringStack := stack.From("Albert Einstein", "Isaac Newton", "Galileo Galilei")

	// Will print from top to bottom:
	// Top element has index = 0
	// 0: Albert Einstein, 1: Isaac Newton, 2: Galileo Galilei
	stringStack.ForEach(func(index int, element string) {
		fmt.Printf("Index: %d, Element: %s\n", index, element)
	})

	// Push to top of the stack
	stringStack.Push("Michael Faraday")
	stringStack.Push("Stephen Hawking")

	// Pop and print out items
	// Stephen Hawking, Michael Faraday, Albert Einstein, Isaac Newton, Galileo Galilei
	for !stringStack.IsEmpty() {
		topElement := stringStack.Pop()
		println(topElement)
	}

	// The stack is empty now
	println(stringStack.IsEmpty()) // true
}

```

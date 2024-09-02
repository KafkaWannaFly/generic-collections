# Queue

`Queue` implements [iindexablecollection.md](iindexablecollection.md "mention") and provides FIFO (First In, First Out) operations. Similar to [stack.md](stack.md "mention"), it has a [linkedlist.md](linkedlist.md "mention") under the hood.

<figure><img src="../.gitbook/assets/9251q1.jpg" alt=""><figcaption></figcaption></figure>

```go
package main

import (
	"fmt"
	"github.com/KafkaWannaFly/generic-collections/queue"
)

func main() {
	stringQueue := queue.From("Go to school", "Go to college", "Get a job", "Get married", "Have kids", "Retire", "Go to heaven")

	// Go from start to end
	// 0: Go to school -> ... -> 6: Go to heaven
	stringQueue.ForEach(func(index int, s string) {
		fmt.Printf("Index %d: %s\n", index, s)
	})

	// Add elements to the end of the queue
	stringQueue.Enqueue("Re-born")
	stringQueue.Enqueue("Go to school again")

	for !stringQueue.IsEmpty() {
		// Remove and return elements from the start of the queue
		firstOut := stringQueue.Dequeue()

		// Print the removed element
		// 0: Go to school -> ... -> 6: Go to heaven -> 7: Re-born -> 8: Go to school again
		fmt.Printf("First out: %s\n", firstOut)
	}

	// The queue is empty now
	fmt.Println("Is the queue empty?", stringQueue.IsEmpty()) // true
}

```

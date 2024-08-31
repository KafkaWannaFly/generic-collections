# LinkedList

`LinkedList` implements [iindexablecollection.md](iindexablecollection.md "mention") like [list.md](list.md "mention"). It has similar methods like `List` with some differences in time complexity.&#x20;

The one thing special about `LinkedList` is its node. And that's it.

```go
package main

import (
	"fmt"
	"github.com/KafkaWannaFly/generic-collections/linkedlist"
)

func main() {
	stringLinkedList := linkedlist.From("Shall Tear", "Demigure", "Cocytus", "Aura", "Mare", "Aureole", "Pandora's Actor", "Albedo")

	node := stringLinkedList.NodeAt(0)
	fmt.Printf("Node at index 0: %v\n", node.Value) // Shall Tear

	secondNode := node.Next
	fmt.Printf("Node at index 1: %v\n", secondNode.Value) // Demigure

	fmt.Printf("Tail node: %v\n", stringLinkedList.Tail.Value) // Albedo
}

```

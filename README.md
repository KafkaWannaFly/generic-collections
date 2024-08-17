---
description: Abstraction prevents code from repeating itself
---

# Quickstart

### Hello, My Old Friend

Generic was not a thing in Golang until v1.18 (Feb 2022). We're implementing different types of data structures which adopt generic feature. People can enjoy type checking and have better developer experience.

```sh
go get github.com/KafkaWannaFly/generic-collections
```

### Working with Struct

This library works well with basic, primitive data-type. If you have any custom struct, which usually yes, you should implement `IHashCoder` and `ILesser` interfaces. They're used by library under the hood. E.g. `Set` would use `IHashCoder` to determine if 2 objects are the same. `List` would use `ILesser` when sorting items.

#### Example

```go
type Book struct {
	Title         string
	Author        string
	CurrentPage   int
	Pages         []string
	PublishedYear int
	Price         float64
}

// A book is define by its title
// If 2 books has same title, they're considered to be equal
func (receiver Book) HashCode() string {
	return receiver.Title
}

// Compare with other book by the title
func (receiver Book) Less(other Book) bool {
	return receiver.Title < other.Title
}
```

#### Interface Definition

{% tabs %}
{% tab title="IHashCoder" %}
```go
package interfaces

type IHashCoder interface {
	// HashCode returns the hash code of the object.
	// Two objects that are equal should have the same hash code.
	HashCode() string
}
```
{% endtab %}

{% tab title="ILesser[TType any]" %}
```go
package interfaces

type ILesser[TType any] interface {
    // Less checks if the current item is less than the given item.
    Less(TType) bool
}
```
{% endtab %}
{% endtabs %}

#### Can I Use Struct Without Implementing the Above Interfaces?

Yes, you can. The system will try to convert your struct into a string to compare.&#x20;

#### **What If My Object Can’t Be Converted?**

GGWP  ╮（╯＿╰）╭

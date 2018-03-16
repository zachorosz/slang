package slang

import "fmt"

// Sequence is an interface for sequential composite types. Sequences are immutable collections that
// are represented by its abstractions.
type Sequence interface {
	Append(items LangType) Sequence
	First() LangType
	Rest() Sequence
	Nth(n Number) LangType
	Len() Number
}

type node struct {
	value LangType
	next  *node
}

// List is a sequence type that is implemented as a singly linked list, Lists have O(1) insertions
// and O(n) access times (if not accessing the head). An empty list, `'()`, is also treated as nil.
type List struct {
	head *node
	tail *node
	len  int
}

// Append - O(1) - returns a new copy of the List with a new item appended at the tail.
func (lst List) Append(obj LangType) Sequence {
	n := &node{
		value: obj,
		next:  nil,
	}

	if lst.head == nil {
		lst.head = n
	} else {
		lst.tail.next = n
	}

	lst.tail = n
	lst.len++

	return lst
}

// First - O(1) - returns the head of the List.
func (lst List) First() LangType {
	return lst.head.value
}

// Rest - O(n) - returns a new List with items starting after the head up to and including the tail.
// If the List has only 1 item, an empty list is returned.
func (lst List) Rest() Sequence {
	rest := List{}
	node := lst.head.next
	for node != nil {
		rest = rest.Append(node.value).(List)
		node = node.next
	}
	return rest
}

// Nth - O(n) - accesses and return the Nth (zero-based) item in the List.
func (lst List) Nth(n Number) LangType {
	node := lst.head
	N := int(n)
	for i := 0; i < N; i++ {
		node = node.next
	}
	return node.value
}

// Len - returns the length of the List.
func (lst List) Len() Number {
	return Number(lst.len)
}

// String - returns a string with the external representation of the List.
func (lst List) String() string {
	items := ""
	node := lst.head
	counter := Number(0)
	for node != nil {
		if counter < lst.Len()-1 {
			items += fmt.Sprintf("%s ", node.value)
		} else {
			items += fmt.Sprint(node.value)
		}
		counter++
		node = node.next
	}
	return fmt.Sprintf("(%s)", items)
}

// Vector is a sequence type that represents a contiguously allocated, dynamic array structure.
// Vectors provide fast accesses in O(1) time and amortized insertions at the tail.
type Vector []LangType

// Append - O(1) amortized - returns a new copy of the Vector with an item appended. If
// re-allocations need to be made, this will not take constant time.
func (vec Vector) Append(obj LangType) Sequence {
	return append(vec, obj)
}

// First - O(1) - returns the first item in the Vector.
func (vec Vector) First() LangType {
	return vec[0]
}

// Rest - O(n) - returns a new Vector with the items starting at index 1 up to and including the
// last item.
func (vec Vector) Rest() Sequence {
	return vec[1:]
}

// Nth - O(1) - accesses and returns the Nth (zero-based) item in the Vector.
func (vec Vector) Nth(n Number) LangType {
	index := int(n)
	return vec[index]
}

// Len returns the length of the Vector.
func (vec Vector) Len() Number {
	return Number(len(vec))
}

func (vec Vector) String() string {
	items := ""
	for i, item := range vec {
		if i < len(vec)-1 {
			items += fmt.Sprintf("%s ", item)
		} else {
			items += fmt.Sprint(item)
		}
	}
	return fmt.Sprintf("[%s]", items)
}

// SequenceP return true if object is a Sequence.
// Usage: `(seq? x)`
func SequenceP(x LangType) bool {
	_, isSeq := x.(Sequence)
	return isSeq
}

// ListP is a predicate that returns true if object is a List.
// Usage: `(list? x)`
func ListP(x LangType) bool {
	x, isList := x.(List)
	return isList
}

// VectorP is a predicate that returns true if object is a Vector.
// Usage: `(vec? x)`
func VectorP(x LangType) bool {
	_, isVec := x.(Vector)
	return isVec
}

// Append returns a new copy of Sequence with appended item(s).
// List append:   O(1)
// Vector append: O(1) amortized
// Usage: `(append seq item)`
func Append(seq Sequence, item LangType) Sequence {
	return seq.Append(item)
}

// Nth returns the nth (zero-based) value of a Sequence.
// List access:   O(n)
// Vector access: O(1)
// Usage: `(nth seq n)`
func Nth(seq Sequence, n Number) (LangType, error) {
	if n < 0 || n >= seq.Len() {
		return nil, fmt.Errorf("Number out of bounds")
	}
	return seq.Nth(n), nil
}

// Len returns the length of a Sequence.
// Usage: `(len seq)`
func Len(seq Sequence) (Number, error) {
	return seq.Len(), nil
}

// MakeList creates a new List from a given set of item(s).
// Usage: `(list items...)`
func MakeList(first LangType, rest ...LangType) List {
	lst := List{}

	lst = lst.Append(first).(List)

	for _, item := range rest {
		lst = lst.Append(item).(List)
	}

	return lst
}

// MakeVector creates a new Vector from a given set of item(s).
// Usage: `(vec items...)`
func MakeVector(first LangType, rest ...LangType) Vector {
	vec := make(Vector, 1+len(rest))

	vec[0] = first

	for i, item := range rest {
		vec[i+1] = item
	}

	return vec
}

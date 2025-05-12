package singelLink

import (
	"fmt"
	"strings"
)

// gen is a type constraint for int, string, or float64
type Gen interface {
	~int | ~string | ~float64
}

type node[T Gen] struct {
	value T
	next  *node[T]
}

type Link[T Gen] struct {
	head  *node[T]
	tail  *node[T]
	count int
}

// Add puts a new value at the head of the list
func (l *Link[T]) Add(val T) {
	myNode := &node[T]{val, l.head}
	l.head = myNode
	if l.tail == nil {
		l.tail = myNode
	}
	l.count++
}

// Insert inserts a new value at a given index
func (l *Link[T]) Insert(val T, idx int) error {
	if idx < 0 || idx >= l.count {
		return fmt.Errorf("index %d is out of bounds for count %d", idx, l.count)
	}

	if idx == 0 {
		l.Add(val)
		return nil
	}

	curNode := l.head
	for i := 0; i < idx-1; i++ {
		curNode = curNode.next
	}

	myNode := &node[T]{val, curNode.next}
	curNode.next = myNode
	if myNode.next == nil {
		l.tail = myNode
	}
	l.count++
	return nil
}

// Get returns the value at the given index
func (l *Link[T]) Get(idx int) (T, error) {
	var zero T
	if idx < 0 || idx >= l.count {
		return zero, fmt.Errorf("index %d is out of bounds for count %d", idx, l.count)
	}

	curNode := l.head
	for i := 0; i < idx; i++ {
		curNode = curNode.next
	}
	return curNode.value, nil
}

// Remove removes and returns the first value in the list
func (l *Link[T]) Remove() (T, error) {
	var zero T
	if l.head == nil {
		return zero, fmt.Errorf("list is empty")
	}
	val := l.head.value
	l.head = l.head.next
	l.count--
	if l.head == nil {
		l.tail = nil
	}
	return val, nil
}

// RemoveLast removes and returns the last value in the list
func (l *Link[T]) RemoveLast() (T, error) {
	var zero T
	if l.head == nil {
		return zero, fmt.Errorf("list is empty")
	}

	if l.head.next == nil {
		val := l.head.value
		l.head = nil
		l.tail = nil
		l.count--
		return val, nil
	}

	curNode := l.head
	for curNode.next.next != nil {
		curNode = curNode.next
	}

	val := curNode.next.value
	curNode.next = nil
	l.tail = curNode
	l.count--
	return val, nil
}

// RemoveAt removes and returns the value at a given index
func (l *Link[T]) RemoveAt(idx int) (T, error) {
	var zero T
	if idx < 0 || idx >= l.count {
		return zero, fmt.Errorf("index %d is out of bounds for count %d", idx, l.count)
	}

	if idx == 0 {
		return l.Remove()
	}

	curNode := l.head
	for i := 0; i < idx-1; i++ {
		curNode = curNode.next
	}

	removedNode := curNode.next
	curNode.next = removedNode.next
	if curNode.next == nil {
		l.tail = curNode
	}
	l.count--
	return removedNode.value, nil
}

// Clear removes all values in the list
func (l *Link[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.count = 0
}

// Search searches for a value and returns the first index or -1 if not found
func (l *Link[T]) Search(val T) int {
	curNode := l.head
	for i := 0; i < l.count; i++ {
		if curNode.value == val {
			return i
		}
		curNode = curNode.next
	}
	return -1
}

// ToString returns a string representation of the list values
func (l *Link[T]) ToString() string {
	var values []string
	current := l.head
	for current != nil {
		values = append(values, fmt.Sprintf("%v", current.value))
		current = current.next
	}
	return strings.Join(values, ", ")
}

// NewSingleLinkList creates and returns a new single linked list
func NewSingleLinkList[T Gen]() Link[T] {
	return Link[T]{nil, nil, 0}
}

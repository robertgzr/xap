package queue

import (
	"bytes"
	"fmt"
)

// Link is a segment of the generic list that hold pointers to the previous and the next linked element
type Link struct {
	next *Link
	prev *Link
	data Track
}

// Next returns the following link
func (lnk *Link) Next() *Link {
	return lnk.next
}

// Prev returns the preceding link
func (lnk *Link) Prev() *Link {
	return lnk.prev
}

// Data returns the data attached to the list segment
func (lnk *Link) Data() Track {
	return lnk.data
}

// List is the data-structure holding a doubly linked list
type List struct {
	first  *Link
	last   *Link
	length int
}

// NewList returns a new empty doubly linked list
func NewList() *List {
	return &List{
		first:  nil,
		last:   nil,
		length: 0,
	}
}

// First returns the first element in the list
func (lst *List) First() Track {
	return lst.first.data
}

// Last returns the last element in the list
func (lst *List) Last() Track {
	return lst.last.data
}

// Len returns the length of the list
func (lst *List) Len() int {
	return lst.length
}

// String returns a string representation of the list
func (lst *List) String() string {
	var buf bytes.Buffer

	var n int
	for lnk := lst.first; lnk != nil; lnk = lnk.next {
		buf.WriteString(fmt.Sprintf("%d : %+v\n", n, lnk.data))
		n++
	}

	return buf.String()
}

// Append takes data and appends it to the end of the list
func (lst *List) Append(data Track) {
	new := Link{nil, nil, data}

	if lst.Len() == 0 {
		lst.first = &new
		lst.last = &new
	} else {
		lst.last.next = &new
		new.prev = lst.last
		lst.last = &new
	}

	lst.length++
}

// Prepend takes data and puts it at the first position in the list
func (lst *List) Prepend(data Track) {
	new := Link{nil, nil, data}

	if lst.Len() == 0 {
		lst.first = &new
		lst.last = &new
	} else {
		lst.first.prev = &new
		new.next = lst.first
		lst.first = &new
	}

	lst.length++
}

// Pop returns and removes the last element int the list
func (lst *List) Pop() Track {
	out := lst.last

	out.prev.next = nil
	lst.last = out.prev

	lst.length--
	return out.data
}

// func (lst *List) Insert(pos int, data Track) {
// }

// func (lst *List) Remove(pos int) {

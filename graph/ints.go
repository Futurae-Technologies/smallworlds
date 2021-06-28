package graph

import (
	"strconv"
)

// IntNode is a wrapper around the int type.
// It provides a simple Node implementation.
type IntNode int

// String converts the underlying int to a string.
func (n IntNode) String() string {
	return strconv.Itoa(int(n))
}

// IntEdge is a slice of IntNodes.
type IntEdge []IntNode

// From returns the first element of the slice.
func (e IntEdge) From() Node {
	return e[0]
}

// To returns the second element of the slice.
func (e IntEdge) To() Node {
	return e[1]
}

// TupleEdge is a slice of Nodes.
type TupleEdge []Node

// From returns the first element.
func (e TupleEdge) From() Node {
	return e[0]
}

// To returns the second element.
func (e TupleEdge) To() Node {
	return e[1]
}

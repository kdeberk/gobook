package treesort

import (
	"bytes"
 "fmt"
)

type tree struct {
	value       int
	left, right *tree
}

func (t *tree) String() string {
	var buf bytes.Buffer

	if nil != t.left {
		buf.WriteString(t.left.String())
		buf.WriteString(", ")
	}

	buf.WriteString(fmt.Sprintf("%d", t.value))
	
	if nil != t.right {
		buf.WriteString(", ")
		buf.WriteString(t.right.String())
	}
	
	return buf.String()
}

// Sort sorts value in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

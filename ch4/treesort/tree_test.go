package treesort

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	xs := []int{9, 8, 3, 2, 1, 4, 5}
	Sort(xs)

	if !reflect.DeepEqual(xs, []int{1, 2, 3, 4, 5, 8, 9}) {
		t.Errorf("Slice not sorted: %v", xs)
	}
}

func TestPrint(t *testing.T) {
	xs := []int{9, 8, 3, 2, 1, 4, 5}

	var root *tree
	for _, v := range xs {
		root = add(root, v)
	}

	if root.String() != "1, 2, 3, 4, 5, 8, 9" {
		t.Errorf("Unexpected string value: %v", root)
	}
}

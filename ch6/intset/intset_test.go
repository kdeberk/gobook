package intset

import (
	"reflect"
	"testing"
)

func TestHas(t *testing.T) {
	var i IntSet
	i.AddAll(2, 100)

	if !i.Has(2) {
		t.Fatalf("expected %v to have element 2", i)
	}
	if i.Has(3) {
		t.Fatalf("expected %v to not have element 3", i)
	}
	if !i.Has(100) {
		t.Fatalf("expected %v to have element 100", i)
	}
}

func TestLen(t *testing.T) {
	var i IntSet
	if 0 != i.Len() {
		t.Fatalf("expected %v to be empty", i)
	}
	i.AddAll(1, 2, 3, 4)
	if 4 != i.Len() {
		t.Fatalf("expected %v to have length 4", i)
	}
}

func TestAdd(t *testing.T) {
	var i IntSet
	if i.Has(2) {
		t.Fatalf("expected %v to not have element 2", i)
	}
	i.Add(2)
	if !i.Has(2) {
		t.Fatalf("expected %v to have element 2", i)
	}
}

func TestUnionWith(t *testing.T) {
	var i, j IntSet
	i.AddAll(1, 2)
	j.AddAll(2, 3, 4)

	i.UnionWith(&j)

	got := i.Elems()
	want := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected %v to have elements %v, got: %v", i, want, got)
	}
}

func TestRemove(t *testing.T) {
	var i IntSet
	i.AddAll(2, 3, 4, 100, 101)
	i.Remove(3)
	i.Remove(100)

	got := i.Elems()
	want := []int{2, 4, 101}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected %v to have elements %v, got: %v", i, want, got)
	}
}

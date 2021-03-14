package main

import (
	"io"
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {
	var r io.Reader
	r = strings.NewReader("Hello, World")
	r = LimitReader(r, 5)

	bs := make([]byte, 6)
	n, err := r.Read(bs)
	switch {
	case n != 5:
		t.Errorf("Expected n = 5, got %v", n)
	case err != nil:
		t.Errorf("Expected no error, got %v", err)
	case string(bs) != "Hello\x00":
		t.Errorf("Expected string, got %v", bs)
	}

	n, err = r.Read(bs)
	switch {
	case n != 0:
		t.Errorf("Expected n = 0, got %v", n)
	case err == nil:
		t.Errorf("Expected error, got nil")
	}
}

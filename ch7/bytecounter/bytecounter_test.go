package bytecounter

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestByteCounterWrite(t *testing.T) {
	ctr := ByteCounter(0)
	fmt.Fprintf(&ctr, "Testing")
	fmt.Fprintf(&ctr, "123")

	if 10 != ctr {
		t.Errorf("Expected 10, instead got: %v", ctr)
	}
}

func TestWordCounterWrite(t *testing.T) {
	ctr := WordCounter(0)
	fmt.Fprintf(&ctr, "one two three")
	fmt.Fprintf(&ctr, "four five")

	if 5 != ctr {
		t.Errorf("Expected 5, instead got: %v", ctr)
	}
}

func TestCounterWrite(t *testing.T) {
	ctr, i := CounterWriter(ioutil.Discard)
	fmt.Fprintf(ctr, "one two three")
	fmt.Fprintf(ctr, "four five")

	if 22 != *i {
		t.Errorf("Expected 5, instead got: %v", *i)
	}
}

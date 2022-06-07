package dprng

import (
	"fmt"
	"testing"
)

func TestRandReader(t *testing.T) {
	testRandReader(t, 0)
	testRandReader(t, 100)
	testRandReader(t, 5000)
}

func testRandReader(t *testing.T, seed int64) {
	r := NewRandReader(seed)
	buf := make([]byte, 128)
	if _, err := r.Read(buf); err != nil {
		t.Fatal(err)
	}
	t.Log(fmt.Sprintf("0x%x", buf))
	if _, err := r.Read(buf); err != nil {
		t.Fatal(err)
	}
	t.Log(fmt.Sprintf("0x%x", buf))
}

package dprng

import (
	"fmt"
	"testing"
)

func TestRandReader(t *testing.T) {
	testRandReader(t, 0, "0x86e02a0174bf3ae5cb5683b220f6756e")
	testRandReader(t, 0, "0x86e02a0174bf3ae5cb5683b220f6756e")
	testRandReader(t, 100, "0xd1f14a9b5b9e2b1240a2231ad7987c7b")
	testRandReader(t, 100, "0xd1f14a9b5b9e2b1240a2231ad7987c7b")
	testRandReader(t, 5000, "0xaaec29002fc64b7ebaa37428ca828b79")
	testRandReader(t, 5000, "0xaaec29002fc64b7ebaa37428ca828b79")
}

func testRandReader(t *testing.T, seed int64, expected string) {
	r := NewRandReader(seed)
	buf := make([]byte, 16)
	if _, err := r.Read(buf); err != nil {
		t.Fatal(err)
	}
	s := fmt.Sprintf("0x%x", buf)
	if expected != s {
		t.Fatal(s)
	}
	if _, err := r.Read(buf); err != nil {
		t.Fatal(err)
	}
	s = fmt.Sprintf("0x%x", buf)
	if expected != s {
		t.Fatal(s)
	}
}

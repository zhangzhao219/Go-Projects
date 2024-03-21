package store

import (
	"testing"
)

func TestStore(t *testing.T) {
	monster1 := Newmonster("a", 1, "abc")
	a := monster1.Store("test1.txt")
	if a {
		t.Logf("true")
	}
}

func TestReStore(t *testing.T) {
	monster1 := Newmonster("a", 1, "abc")
	b := monster1.ReStore("test1.txt")
	t.Logf(b.Name)
}

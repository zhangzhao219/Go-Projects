package main

import (
	"testing"
)

func TestAddUpper(t *testing.T) {
	res := addUpper(9)
	if res != 55 {
		t.Fatalf("error!")
	}
	t.Logf("Right!")
}

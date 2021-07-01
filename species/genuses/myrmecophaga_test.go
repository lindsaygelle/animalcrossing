package genuses

import "testing"

func TestMyrmecophaga(t *testing.T) {
	var s string = "Myrmecophaga"
	if ok := myrmecophaga == s; !ok {
		t.Fatalf("myrmecophaga != %s", s)
	}
}

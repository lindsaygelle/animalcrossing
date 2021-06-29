package genuses

import "testing"

func TestGenusMeleagris(t *testing.T) {
	var s string = "Meleagris"
	if ok := meleagris == s; !ok {
		t.Fatalf("meleagris != %s", s)
	}
}

package genuses

import "testing"

func TestMeleagris(t *testing.T) {
	var s string = "Meleagris"
	if ok := meleagris == s; !ok {
		t.Fatalf("meleagris != %s", s)
	}
}

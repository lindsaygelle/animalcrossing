package animal

import "testing"

// TestBeaverId tests whether Beaver has the correct ID.
func TestBeaverId(t *testing.T) {
	if v := Beaver.Id(); !(v == beaverId) {
		t.Fatalf("%s != %s", v, beaverId)
	}	
}

// TestBeaverFictional test whether Beaver is a fictional Animal Crossing animal type.
func TestBeaverFictional(t *testing.T) {
	if v := Beaver.Fictional(); !(v == beaverFictional) {
		t.Fatalf("%t != %t", v, beaverFictional)
	}
}

// TestBeaverSpecial tests whether Beaver is a special Animal Crossing animal type.
func TestBeaverSpecial(t *testing.T) {
	if v := Beaver.Special(); !(v == beaverSpecial) {
		t.Fatalf("%t != %t", v, beaverSpecial)
	}
}

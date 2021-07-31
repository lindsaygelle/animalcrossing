package animal

import "testing"

// TestElephantId tests whether Elephant has the correct ID.
func TestElephantId(t *testing.T) {
	if v := Elephant.Id(); !(v == elephantId) {
		t.Fatalf("%s != %s", v, elephantId)
	}	
}

// TestElephantFictional test whether Elephant is a fictional Animal Crossing animal type.
func TestElephantFictional(t *testing.T) {
	if v := Elephant.Fictional(); !(v == elephantFictional) {
		t.Fatalf("%t != %t", v, elephantFictional)
	}
}

// TestElephantSpecial tests whether Elephant is a special Animal Crossing animal type.
func TestElephantSpecial(t *testing.T) {
	if v := Elephant.Special(); !(v == elephantSpecial) {
		t.Fatalf("%t != %t", v, elephantSpecial)
	}
}

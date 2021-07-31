package animal

import "testing"

// TestAlligatorId tests whether Alligator has the correct ID.
func TestAlligatorId(t *testing.T) {
	if v := Alligator.Id(); !(v == alligatorId) {
		t.Fatalf("%s != %s", v, alligatorId)
	}	
}

// TestAlligatorFictional test whether Alligator is a fictional Animal Crossing animal type.
func TestAlligatorFictional(t *testing.T) {
	if v := Alligator.Fictional(); !(v == alligatorFictional) {
		t.Fatalf("%t != %t", v, alligatorFictional)
	}
}

// TestAlligatorSpecial tests whether Alligator is a special Animal Crossing animal type.
func TestAlligatorSpecial(t *testing.T) {
	if v := Alligator.Special(); !(v == alligatorSpecial) {
		t.Fatalf("%t != %t", v, alligatorSpecial)
	}
}

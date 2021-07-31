package animal

import "testing"

// TestHamsterId tests whether Hamster has the correct ID.
func TestHamsterId(t *testing.T) {
	if v := Hamster.Id(); !(v == hamsterId) {
		t.Fatalf("%s != %s", v, hamsterId)
	}	
}

// TestHamsterFictional test whether Hamster is a fictional Animal Crossing animal type.
func TestHamsterFictional(t *testing.T) {
	if v := Hamster.Fictional(); !(v == hamsterFictional) {
		t.Fatalf("%t != %t", v, hamsterFictional)
	}
}

// TestHamsterSpecial tests whether Hamster is a special Animal Crossing animal type.
func TestHamsterSpecial(t *testing.T) {
	if v := Hamster.Special(); !(v == hamsterSpecial) {
		t.Fatalf("%t != %t", v, hamsterSpecial)
	}
}

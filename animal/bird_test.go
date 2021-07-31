package animal

import "testing"

// TestBirdId tests whether Bird has the correct ID.
func TestBirdId(t *testing.T) {
	if v := Bird.Id(); !(v == birdId) {
		t.Fatalf("%s != %s", v, birdId)
	}	
}

// TestBirdFictional test whether Bird is a fictional Animal Crossing animal type.
func TestBirdFictional(t *testing.T) {
	if v := Bird.Fictional(); !(v == birdFictional) {
		t.Fatalf("%t != %t", v, birdFictional)
	}
}

// TestBirdSpecial tests whether Bird is a special Animal Crossing animal type.
func TestBirdSpecial(t *testing.T) {
	if v := Bird.Special(); !(v == birdSpecial) {
		t.Fatalf("%t != %t", v, birdSpecial)
	}
}

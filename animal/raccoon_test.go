package animal

import "testing"

// TestRaccoonId tests whether Raccoon has the correct ID.
func TestRaccoonId(t *testing.T) {
	if v := Raccoon.Id(); !(v == raccoonId) {
		t.Fatalf("%s != %s", v, raccoonId)
	}	
}

// TestRaccoonFictional test whether Raccoon is a fictional Animal Crossing animal type.
func TestRaccoonFictional(t *testing.T) {
	if v := Raccoon.Fictional(); !(v == raccoonFictional) {
		t.Fatalf("%t != %t", v, raccoonFictional)
	}
}

// TestRaccoonSpecial tests whether Raccoon is a special Animal Crossing animal type.
func TestRaccoonSpecial(t *testing.T) {
	if v := Raccoon.Special(); !(v == raccoonSpecial) {
		t.Fatalf("%t != %t", v, raccoonSpecial)
	}
}

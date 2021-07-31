package animal

import "testing"

// TestGoatId tests whether Goat has the correct ID.
func TestGoatId(t *testing.T) {
	if v := Goat.Id(); !(v == goatId) {
		t.Fatalf("%s != %s", v, goatId)
	}	
}

// TestGoatFictional test whether Goat is a fictional Animal Crossing animal type.
func TestGoatFictional(t *testing.T) {
	if v := Goat.Fictional(); !(v == goatFictional) {
		t.Fatalf("%t != %t", v, goatFictional)
	}
}

// TestGoatSpecial tests whether Goat is a special Animal Crossing animal type.
func TestGoatSpecial(t *testing.T) {
	if v := Goat.Special(); !(v == goatSpecial) {
		t.Fatalf("%t != %t", v, goatSpecial)
	}
}

package animal

import "testing"

// TestFrogId tests whether Frog has the correct ID.
func TestFrogId(t *testing.T) {
	if v := Frog.Id(); !(v == frogId) {
		t.Fatalf("%s != %s", v, frogId)
	}	
}

// TestFrogFictional test whether Frog is a fictional Animal Crossing animal type.
func TestFrogFictional(t *testing.T) {
	if v := Frog.Fictional(); !(v == frogFictional) {
		t.Fatalf("%t != %t", v, frogFictional)
	}
}

// TestFrogSpecial tests whether Frog is a special Animal Crossing animal type.
func TestFrogSpecial(t *testing.T) {
	if v := Frog.Special(); !(v == frogSpecial) {
		t.Fatalf("%t != %t", v, frogSpecial)
	}
}

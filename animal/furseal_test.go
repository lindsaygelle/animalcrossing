package animal

import "testing"

// TestFursealId tests whether Furseal has the correct ID.
func TestFursealId(t *testing.T) {
	if v := Furseal.Id(); !(v == fursealId) {
		t.Fatalf("%s != %s", v, fursealId)
	}	
}

// TestFursealFictional test whether Furseal is a fictional Animal Crossing animal type.
func TestFursealFictional(t *testing.T) {
	if v := Furseal.Fictional(); !(v == fursealFictional) {
		t.Fatalf("%t != %t", v, fursealFictional)
	}
}

// TestFursealSpecial tests whether Furseal is a special Animal Crossing animal type.
func TestFursealSpecial(t *testing.T) {
	if v := Furseal.Special(); !(v == fursealSpecial) {
		t.Fatalf("%t != %t", v, fursealSpecial)
	}
}

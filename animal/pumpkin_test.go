package animal

import "testing"

// TestPumpkinId tests whether Pumpkin has the correct ID.
func TestPumpkinId(t *testing.T) {
	if v := Pumpkin.Id(); !(v == pumpkinId) {
		t.Fatalf("%s != %s", v, pumpkinId)
	}	
}

// TestPumpkinFictional test whether Pumpkin is a fictional Animal Crossing animal type.
func TestPumpkinFictional(t *testing.T) {
	if v := Pumpkin.Fictional(); !(v == pumpkinFictional) {
		t.Fatalf("%t != %t", v, pumpkinFictional)
	}
}

// TestPumpkinSpecial tests whether Pumpkin is a special Animal Crossing animal type.
func TestPumpkinSpecial(t *testing.T) {
	if v := Pumpkin.Special(); !(v == pumpkinSpecial) {
		t.Fatalf("%t != %t", v, pumpkinSpecial)
	}
}

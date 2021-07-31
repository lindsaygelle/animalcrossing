package animal

import "testing"

// TestSlothId tests whether Sloth has the correct ID.
func TestSlothId(t *testing.T) {
	if v := Sloth.Id(); !(v == slothId) {
		t.Fatalf("%s != %s", v, slothId)
	}	
}

// TestSlothFictional test whether Sloth is a fictional Animal Crossing animal type.
func TestSlothFictional(t *testing.T) {
	if v := Sloth.Fictional(); !(v == slothFictional) {
		t.Fatalf("%t != %t", v, slothFictional)
	}
}

// TestSlothSpecial tests whether Sloth is a special Animal Crossing animal type.
func TestSlothSpecial(t *testing.T) {
	if v := Sloth.Special(); !(v == slothSpecial) {
		t.Fatalf("%t != %t", v, slothSpecial)
	}
}

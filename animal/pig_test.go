package animal

import "testing"

// TestPigId tests whether Pig has the correct ID.
func TestPigId(t *testing.T) {
	if v := Pig.Id(); !(v == pigId) {
		t.Fatalf("%s != %s", v, pigId)
	}	
}

// TestPigFictional test whether Pig is a fictional Animal Crossing animal type.
func TestPigFictional(t *testing.T) {
	if v := Pig.Fictional(); !(v == pigFictional) {
		t.Fatalf("%t != %t", v, pigFictional)
	}
}

// TestPigSpecial tests whether Pig is a special Animal Crossing animal type.
func TestPigSpecial(t *testing.T) {
	if v := Pig.Special(); !(v == pigSpecial) {
		t.Fatalf("%t != %t", v, pigSpecial)
	}
}

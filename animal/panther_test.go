package animal

import "testing"

// TestPantherId tests whether Panther has the correct ID.
func TestPantherId(t *testing.T) {
	if v := Panther.Id(); !(v == pantherId) {
		t.Fatalf("%s != %s", v, pantherId)
	}	
}

// TestPantherFictional test whether Panther is a fictional Animal Crossing animal type.
func TestPantherFictional(t *testing.T) {
	if v := Panther.Fictional(); !(v == pantherFictional) {
		t.Fatalf("%t != %t", v, pantherFictional)
	}
}

// TestPantherSpecial tests whether Panther is a special Animal Crossing animal type.
func TestPantherSpecial(t *testing.T) {
	if v := Panther.Special(); !(v == pantherSpecial) {
		t.Fatalf("%t != %t", v, pantherSpecial)
	}
}

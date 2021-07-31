package animal

import "testing"

// TestGorillaId tests whether Gorilla has the correct ID.
func TestGorillaId(t *testing.T) {
	if v := Gorilla.Id(); !(v == gorillaId) {
		t.Fatalf("%s != %s", v, gorillaId)
	}	
}

// TestGorillaFictional test whether Gorilla is a fictional Animal Crossing animal type.
func TestGorillaFictional(t *testing.T) {
	if v := Gorilla.Fictional(); !(v == gorillaFictional) {
		t.Fatalf("%t != %t", v, gorillaFictional)
	}
}

// TestGorillaSpecial tests whether Gorilla is a special Animal Crossing animal type.
func TestGorillaSpecial(t *testing.T) {
	if v := Gorilla.Special(); !(v == gorillaSpecial) {
		t.Fatalf("%t != %t", v, gorillaSpecial)
	}
}

package animal

import "testing"

// TestDeerId tests whether Deer has the correct ID.
func TestDeerId(t *testing.T) {
	if v := Deer.Id(); !(v == deerId) {
		t.Fatalf("%s != %s", v, deerId)
	}	
}

// TestDeerFictional test whether Deer is a fictional Animal Crossing animal type.
func TestDeerFictional(t *testing.T) {
	if v := Deer.Fictional(); !(v == deerFictional) {
		t.Fatalf("%t != %t", v, deerFictional)
	}
}

// TestDeerSpecial tests whether Deer is a special Animal Crossing animal type.
func TestDeerSpecial(t *testing.T) {
	if v := Deer.Special(); !(v == deerSpecial) {
		t.Fatalf("%t != %t", v, deerSpecial)
	}
}

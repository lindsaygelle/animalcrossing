package animal

import "testing"

// TestDodoId tests whether Dodo has the correct ID.
func TestDodoId(t *testing.T) {
	if v := Dodo.Id(); !(v == dodoId) {
		t.Fatalf("%s != %s", v, dodoId)
	}	
}

// TestDodoFictional test whether Dodo is a fictional Animal Crossing animal type.
func TestDodoFictional(t *testing.T) {
	if v := Dodo.Fictional(); !(v == dodoFictional) {
		t.Fatalf("%t != %t", v, dodoFictional)
	}
}

// TestDodoSpecial tests whether Dodo is a special Animal Crossing animal type.
func TestDodoSpecial(t *testing.T) {
	if v := Dodo.Special(); !(v == dodoSpecial) {
		t.Fatalf("%t != %t", v, dodoSpecial)
	}
}

package animal

import "testing"

// TestBearId tests whether Bear has the correct ID.
func TestBearId(t *testing.T) {
	if v := Bear.Id(); !(v == bearId) {
		t.Fatalf("%s != %s", v, bearId)
	}	
}

// TestBearFictional test whether Bear is a fictional Animal Crossing animal type.
func TestBearFictional(t *testing.T) {
	if v := Bear.Fictional(); !(v == bearFictional) {
		t.Fatalf("%t != %t", v, bearFictional)
	}
}

// TestBearSpecial tests whether Bear is a special Animal Crossing animal type.
func TestBearSpecial(t *testing.T) {
	if v := Bear.Special(); !(v == bearSpecial) {
		t.Fatalf("%t != %t", v, bearSpecial)
	}
}

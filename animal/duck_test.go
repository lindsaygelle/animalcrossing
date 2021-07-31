package animal

import "testing"

// TestDuckId tests whether Duck has the correct ID.
func TestDuckId(t *testing.T) {
	if v := Duck.Id(); !(v == duckId) {
		t.Fatalf("%s != %s", v, duckId)
	}	
}

// TestDuckFictional test whether Duck is a fictional Animal Crossing animal type.
func TestDuckFictional(t *testing.T) {
	if v := Duck.Fictional(); !(v == duckFictional) {
		t.Fatalf("%t != %t", v, duckFictional)
	}
}

// TestDuckSpecial tests whether Duck is a special Animal Crossing animal type.
func TestDuckSpecial(t *testing.T) {
	if v := Duck.Special(); !(v == duckSpecial) {
		t.Fatalf("%t != %t", v, duckSpecial)
	}
}

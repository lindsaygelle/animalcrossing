package animal

import "testing"

// TestGyroidId tests whether Gyroid has the correct ID.
func TestGyroidId(t *testing.T) {
	if v := Gyroid.Id(); !(v == gyroidId) {
		t.Fatalf("%s != %s", v, gyroidId)
	}	
}

// TestGyroidFictional test whether Gyroid is a fictional Animal Crossing animal type.
func TestGyroidFictional(t *testing.T) {
	if v := Gyroid.Fictional(); !(v == gyroidFictional) {
		t.Fatalf("%t != %t", v, gyroidFictional)
	}
}

// TestGyroidSpecial tests whether Gyroid is a special Animal Crossing animal type.
func TestGyroidSpecial(t *testing.T) {
	if v := Gyroid.Special(); !(v == gyroidSpecial) {
		t.Fatalf("%t != %t", v, gyroidSpecial)
	}
}

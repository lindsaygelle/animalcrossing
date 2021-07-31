package animal

import "testing"

// TestCowId tests whether Cow has the correct ID.
func TestCowId(t *testing.T) {
	if v := Cow.Id(); !(v == cowId) {
		t.Fatalf("%s != %s", v, cowId)
	}	
}

// TestCowFictional test whether Cow is a fictional Animal Crossing animal type.
func TestCowFictional(t *testing.T) {
	if v := Cow.Fictional(); !(v == cowFictional) {
		t.Fatalf("%t != %t", v, cowFictional)
	}
}

// TestCowSpecial tests whether Cow is a special Animal Crossing animal type.
func TestCowSpecial(t *testing.T) {
	if v := Cow.Special(); !(v == cowSpecial) {
		t.Fatalf("%t != %t", v, cowSpecial)
	}
}

package animal

import "testing"

// TestBearcubId tests whether Bearcub has the correct ID.
func TestBearcubId(t *testing.T) {
	if v := Bearcub.Id(); !(v == bearcubId) {
		t.Fatalf("%s != %s", v, bearcubId)
	}	
}

// TestBearcubFictional test whether Bearcub is a fictional Animal Crossing animal type.
func TestBearcubFictional(t *testing.T) {
	if v := Bearcub.Fictional(); !(v == bearcubFictional) {
		t.Fatalf("%t != %t", v, bearcubFictional)
	}
}

// TestBearcubSpecial tests whether Bearcub is a special Animal Crossing animal type.
func TestBearcubSpecial(t *testing.T) {
	if v := Bearcub.Special(); !(v == bearcubSpecial) {
		t.Fatalf("%t != %t", v, bearcubSpecial)
	}
}

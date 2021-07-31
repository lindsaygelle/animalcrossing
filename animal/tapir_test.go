package animal

import "testing"

// TestTapirId tests whether Tapir has the correct ID.
func TestTapirId(t *testing.T) {
	if v := Tapir.Id(); !(v == tapirId) {
		t.Fatalf("%s != %s", v, tapirId)
	}	
}

// TestTapirFictional test whether Tapir is a fictional Animal Crossing animal type.
func TestTapirFictional(t *testing.T) {
	if v := Tapir.Fictional(); !(v == tapirFictional) {
		t.Fatalf("%t != %t", v, tapirFictional)
	}
}

// TestTapirSpecial tests whether Tapir is a special Animal Crossing animal type.
func TestTapirSpecial(t *testing.T) {
	if v := Tapir.Special(); !(v == tapirSpecial) {
		t.Fatalf("%t != %t", v, tapirSpecial)
	}
}

package animal

import "testing"

// TestLionId tests whether Lion has the correct ID.
func TestLionId(t *testing.T) {
	if v := Lion.Id(); !(v == lionId) {
		t.Fatalf("%s != %s", v, lionId)
	}	
}

// TestLionFictional test whether Lion is a fictional Animal Crossing animal type.
func TestLionFictional(t *testing.T) {
	if v := Lion.Fictional(); !(v == lionFictional) {
		t.Fatalf("%t != %t", v, lionFictional)
	}
}

// TestLionSpecial tests whether Lion is a special Animal Crossing animal type.
func TestLionSpecial(t *testing.T) {
	if v := Lion.Special(); !(v == lionSpecial) {
		t.Fatalf("%t != %t", v, lionSpecial)
	}
}

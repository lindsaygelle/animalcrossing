package animal

import "testing"

// TestPeacockId tests whether Peacock has the correct ID.
func TestPeacockId(t *testing.T) {
	if v := Peacock.Id(); !(v == peacockId) {
		t.Fatalf("%s != %s", v, peacockId)
	}	
}

// TestPeacockFictional test whether Peacock is a fictional Animal Crossing animal type.
func TestPeacockFictional(t *testing.T) {
	if v := Peacock.Fictional(); !(v == peacockFictional) {
		t.Fatalf("%t != %t", v, peacockFictional)
	}
}

// TestPeacockSpecial tests whether Peacock is a special Animal Crossing animal type.
func TestPeacockSpecial(t *testing.T) {
	if v := Peacock.Special(); !(v == peacockSpecial) {
		t.Fatalf("%t != %t", v, peacockSpecial)
	}
}

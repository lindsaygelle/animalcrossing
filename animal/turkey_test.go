package animal

import "testing"

// TestTurkeyId tests whether Turkey has the correct ID.
func TestTurkeyId(t *testing.T) {
	if v := Turkey.Id(); !(v == turkeyId) {
		t.Fatalf("%s != %s", v, turkeyId)
	}	
}

// TestTurkeyFictional test whether Turkey is a fictional Animal Crossing animal type.
func TestTurkeyFictional(t *testing.T) {
	if v := Turkey.Fictional(); !(v == turkeyFictional) {
		t.Fatalf("%t != %t", v, turkeyFictional)
	}
}

// TestTurkeySpecial tests whether Turkey is a special Animal Crossing animal type.
func TestTurkeySpecial(t *testing.T) {
	if v := Turkey.Special(); !(v == turkeySpecial) {
		t.Fatalf("%t != %t", v, turkeySpecial)
	}
}

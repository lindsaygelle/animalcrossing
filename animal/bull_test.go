package animal

import "testing"

// TestBullId tests whether Bull has the correct ID.
func TestBullId(t *testing.T) {
	if v := Bull.Id(); !(v == bullId) {
		t.Fatalf("%s != %s", v, bullId)
	}	
}

// TestBullFictional test whether Bull is a fictional Animal Crossing animal type.
func TestBullFictional(t *testing.T) {
	if v := Bull.Fictional(); !(v == bullFictional) {
		t.Fatalf("%t != %t", v, bullFictional)
	}
}

// TestBullSpecial tests whether Bull is a special Animal Crossing animal type.
func TestBullSpecial(t *testing.T) {
	if v := Bull.Special(); !(v == bullSpecial) {
		t.Fatalf("%t != %t", v, bullSpecial)
	}
}

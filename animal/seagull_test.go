package animal

import "testing"

// TestSeagullId tests whether Seagull has the correct ID.
func TestSeagullId(t *testing.T) {
	if v := Seagull.Id(); !(v == seagullId) {
		t.Fatalf("%s != %s", v, seagullId)
	}	
}

// TestSeagullFictional test whether Seagull is a fictional Animal Crossing animal type.
func TestSeagullFictional(t *testing.T) {
	if v := Seagull.Fictional(); !(v == seagullFictional) {
		t.Fatalf("%t != %t", v, seagullFictional)
	}
}

// TestSeagullSpecial tests whether Seagull is a special Animal Crossing animal type.
func TestSeagullSpecial(t *testing.T) {
	if v := Seagull.Special(); !(v == seagullSpecial) {
		t.Fatalf("%t != %t", v, seagullSpecial)
	}
}

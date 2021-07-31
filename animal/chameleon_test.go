package animal

import "testing"

// TestChameleonId tests whether Chameleon has the correct ID.
func TestChameleonId(t *testing.T) {
	if v := Chameleon.Id(); !(v == chameleonId) {
		t.Fatalf("%s != %s", v, chameleonId)
	}	
}

// TestChameleonFictional test whether Chameleon is a fictional Animal Crossing animal type.
func TestChameleonFictional(t *testing.T) {
	if v := Chameleon.Fictional(); !(v == chameleonFictional) {
		t.Fatalf("%t != %t", v, chameleonFictional)
	}
}

// TestChameleonSpecial tests whether Chameleon is a special Animal Crossing animal type.
func TestChameleonSpecial(t *testing.T) {
	if v := Chameleon.Special(); !(v == chameleonSpecial) {
		t.Fatalf("%t != %t", v, chameleonSpecial)
	}
}

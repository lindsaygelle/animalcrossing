package animal

import "testing"

// TestMoleId tests whether Mole has the correct ID.
func TestMoleId(t *testing.T) {
	if v := Mole.Id(); !(v == moleId) {
		t.Fatalf("%s != %s", v, moleId)
	}	
}

// TestMoleFictional test whether Mole is a fictional Animal Crossing animal type.
func TestMoleFictional(t *testing.T) {
	if v := Mole.Fictional(); !(v == moleFictional) {
		t.Fatalf("%t != %t", v, moleFictional)
	}
}

// TestMoleSpecial tests whether Mole is a special Animal Crossing animal type.
func TestMoleSpecial(t *testing.T) {
	if v := Mole.Special(); !(v == moleSpecial) {
		t.Fatalf("%t != %t", v, moleSpecial)
	}
}

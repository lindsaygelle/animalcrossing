package animal

import "testing"

// TestAxolotlId tests whether Axolotl has the correct ID.
func TestAxolotlId(t *testing.T) {
	if v := Axolotl.Id(); !(v == axolotlId) {
		t.Fatalf("%s != %s", v, axolotlId)
	}	
}

// TestAxolotlFictional test whether Axolotl is a fictional Animal Crossing animal type.
func TestAxolotlFictional(t *testing.T) {
	if v := Axolotl.Fictional(); !(v == axolotlFictional) {
		t.Fatalf("%t != %t", v, axolotlFictional)
	}
}

// TestAxolotlSpecial tests whether Axolotl is a special Animal Crossing animal type.
func TestAxolotlSpecial(t *testing.T) {
	if v := Axolotl.Special(); !(v == axolotlSpecial) {
		t.Fatalf("%t != %t", v, axolotlSpecial)
	}
}

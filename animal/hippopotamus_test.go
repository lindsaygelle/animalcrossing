package animal

import "testing"

// TestHippopotamusId tests whether Hippopotamus has the correct ID.
func TestHippopotamusId(t *testing.T) {
	if v := Hippopotamus.Id(); !(v == hippopotamusId) {
		t.Fatalf("%s != %s", v, hippopotamusId)
	}	
}

// TestHippopotamusFictional test whether Hippopotamus is a fictional Animal Crossing animal type.
func TestHippopotamusFictional(t *testing.T) {
	if v := Hippopotamus.Fictional(); !(v == hippopotamusFictional) {
		t.Fatalf("%t != %t", v, hippopotamusFictional)
	}
}

// TestHippopotamusSpecial tests whether Hippopotamus is a special Animal Crossing animal type.
func TestHippopotamusSpecial(t *testing.T) {
	if v := Hippopotamus.Special(); !(v == hippopotamusSpecial) {
		t.Fatalf("%t != %t", v, hippopotamusSpecial)
	}
}

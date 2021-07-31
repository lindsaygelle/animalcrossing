package animal

import "testing"

// TestOctopusId tests whether Octopus has the correct ID.
func TestOctopusId(t *testing.T) {
	if v := Octopus.Id(); !(v == octopusId) {
		t.Fatalf("%s != %s", v, octopusId)
	}	
}

// TestOctopusFictional test whether Octopus is a fictional Animal Crossing animal type.
func TestOctopusFictional(t *testing.T) {
	if v := Octopus.Fictional(); !(v == octopusFictional) {
		t.Fatalf("%t != %t", v, octopusFictional)
	}
}

// TestOctopusSpecial tests whether Octopus is a special Animal Crossing animal type.
func TestOctopusSpecial(t *testing.T) {
	if v := Octopus.Special(); !(v == octopusSpecial) {
		t.Fatalf("%t != %t", v, octopusSpecial)
	}
}

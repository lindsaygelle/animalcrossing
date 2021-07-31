package animal

import "testing"

// TestHorseId tests whether Horse has the correct ID.
func TestHorseId(t *testing.T) {
	if v := Horse.Id(); !(v == horseId) {
		t.Fatalf("%s != %s", v, horseId)
	}	
}

// TestHorseFictional test whether Horse is a fictional Animal Crossing animal type.
func TestHorseFictional(t *testing.T) {
	if v := Horse.Fictional(); !(v == horseFictional) {
		t.Fatalf("%t != %t", v, horseFictional)
	}
}

// TestHorseSpecial tests whether Horse is a special Animal Crossing animal type.
func TestHorseSpecial(t *testing.T) {
	if v := Horse.Special(); !(v == horseSpecial) {
		t.Fatalf("%t != %t", v, horseSpecial)
	}
}

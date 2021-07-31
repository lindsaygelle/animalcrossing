package animal

import "testing"

// TestAnteaterId tests whether Anteater has the correct ID.
func TestAnteaterId(t *testing.T) {
	if v := Anteater.Id(); !(v == anteaterId) {
		t.Fatalf("%s != %s", v, anteaterId)
	}	
}

// TestAnteaterFictional test whether Anteater is a fictional Animal Crossing animal type.
func TestAnteaterFictional(t *testing.T) {
	if v := Anteater.Fictional(); !(v == anteaterFictional) {
		t.Fatalf("%t != %t", v, anteaterFictional)
	}
}

// TestAnteaterSpecial tests whether Anteater is a special Animal Crossing animal type.
func TestAnteaterSpecial(t *testing.T) {
	if v := Anteater.Special(); !(v == anteaterSpecial) {
		t.Fatalf("%t != %t", v, anteaterSpecial)
	}
}

package animal

import "testing"

// TestPenguinId tests whether Penguin has the correct ID.
func TestPenguinId(t *testing.T) {
	if v := Penguin.Id(); !(v == penguinId) {
		t.Fatalf("%s != %s", v, penguinId)
	}	
}

// TestPenguinFictional test whether Penguin is a fictional Animal Crossing animal type.
func TestPenguinFictional(t *testing.T) {
	if v := Penguin.Fictional(); !(v == penguinFictional) {
		t.Fatalf("%t != %t", v, penguinFictional)
	}
}

// TestPenguinSpecial tests whether Penguin is a special Animal Crossing animal type.
func TestPenguinSpecial(t *testing.T) {
	if v := Penguin.Special(); !(v == penguinSpecial) {
		t.Fatalf("%t != %t", v, penguinSpecial)
	}
}

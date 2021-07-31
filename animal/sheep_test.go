package animal

import "testing"

// TestSheepId tests whether Sheep has the correct ID.
func TestSheepId(t *testing.T) {
	if v := Sheep.Id(); !(v == sheepId) {
		t.Fatalf("%s != %s", v, sheepId)
	}	
}

// TestSheepFictional test whether Sheep is a fictional Animal Crossing animal type.
func TestSheepFictional(t *testing.T) {
	if v := Sheep.Fictional(); !(v == sheepFictional) {
		t.Fatalf("%t != %t", v, sheepFictional)
	}
}

// TestSheepSpecial tests whether Sheep is a special Animal Crossing animal type.
func TestSheepSpecial(t *testing.T) {
	if v := Sheep.Special(); !(v == sheepSpecial) {
		t.Fatalf("%t != %t", v, sheepSpecial)
	}
}

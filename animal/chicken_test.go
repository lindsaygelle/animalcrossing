package animal

import "testing"

// TestChickenId tests whether Chicken has the correct ID.
func TestChickenId(t *testing.T) {
	if v := Chicken.Id(); !(v == chickenId) {
		t.Fatalf("%s != %s", v, chickenId)
	}	
}

// TestChickenFictional test whether Chicken is a fictional Animal Crossing animal type.
func TestChickenFictional(t *testing.T) {
	if v := Chicken.Fictional(); !(v == chickenFictional) {
		t.Fatalf("%t != %t", v, chickenFictional)
	}
}

// TestChickenSpecial tests whether Chicken is a special Animal Crossing animal type.
func TestChickenSpecial(t *testing.T) {
	if v := Chicken.Special(); !(v == chickenSpecial) {
		t.Fatalf("%t != %t", v, chickenSpecial)
	}
}

package animal

import "testing"

// TestMonkeyId tests whether Monkey has the correct ID.
func TestMonkeyId(t *testing.T) {
	if v := Monkey.Id(); !(v == monkeyId) {
		t.Fatalf("%s != %s", v, monkeyId)
	}	
}

// TestMonkeyFictional test whether Monkey is a fictional Animal Crossing animal type.
func TestMonkeyFictional(t *testing.T) {
	if v := Monkey.Fictional(); !(v == monkeyFictional) {
		t.Fatalf("%t != %t", v, monkeyFictional)
	}
}

// TestMonkeySpecial tests whether Monkey is a special Animal Crossing animal type.
func TestMonkeySpecial(t *testing.T) {
	if v := Monkey.Special(); !(v == monkeySpecial) {
		t.Fatalf("%t != %t", v, monkeySpecial)
	}
}

package animal

import "testing"

// TestWalrusId tests whether Walrus has the correct ID.
func TestWalrusId(t *testing.T) {
	if v := Walrus.Id(); !(v == walrusId) {
		t.Fatalf("%s != %s", v, walrusId)
	}	
}

// TestWalrusFictional test whether Walrus is a fictional Animal Crossing animal type.
func TestWalrusFictional(t *testing.T) {
	if v := Walrus.Fictional(); !(v == walrusFictional) {
		t.Fatalf("%t != %t", v, walrusFictional)
	}
}

// TestWalrusSpecial tests whether Walrus is a special Animal Crossing animal type.
func TestWalrusSpecial(t *testing.T) {
	if v := Walrus.Special(); !(v == walrusSpecial) {
		t.Fatalf("%t != %t", v, walrusSpecial)
	}
}

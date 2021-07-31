package animal

import "testing"

// TestHedgehogId tests whether Hedgehog has the correct ID.
func TestHedgehogId(t *testing.T) {
	if v := Hedgehog.Id(); !(v == hedgehogId) {
		t.Fatalf("%s != %s", v, hedgehogId)
	}	
}

// TestHedgehogFictional test whether Hedgehog is a fictional Animal Crossing animal type.
func TestHedgehogFictional(t *testing.T) {
	if v := Hedgehog.Fictional(); !(v == hedgehogFictional) {
		t.Fatalf("%t != %t", v, hedgehogFictional)
	}
}

// TestHedgehogSpecial tests whether Hedgehog is a special Animal Crossing animal type.
func TestHedgehogSpecial(t *testing.T) {
	if v := Hedgehog.Special(); !(v == hedgehogSpecial) {
		t.Fatalf("%t != %t", v, hedgehogSpecial)
	}
}

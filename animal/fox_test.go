package animal

import "testing"

// TestFoxId tests whether Fox has the correct ID.
func TestFoxId(t *testing.T) {
	if v := Fox.Id(); !(v == foxId) {
		t.Fatalf("%s != %s", v, foxId)
	}	
}

// TestFoxFictional test whether Fox is a fictional Animal Crossing animal type.
func TestFoxFictional(t *testing.T) {
	if v := Fox.Fictional(); !(v == foxFictional) {
		t.Fatalf("%t != %t", v, foxFictional)
	}
}

// TestFoxSpecial tests whether Fox is a special Animal Crossing animal type.
func TestFoxSpecial(t *testing.T) {
	if v := Fox.Special(); !(v == foxSpecial) {
		t.Fatalf("%t != %t", v, foxSpecial)
	}
}

package animal

import "testing"

// TestSquirrelId tests whether Squirrel has the correct ID.
func TestSquirrelId(t *testing.T) {
	if v := Squirrel.Id(); !(v == squirrelId) {
		t.Fatalf("%s != %s", v, squirrelId)
	}	
}

// TestSquirrelFictional test whether Squirrel is a fictional Animal Crossing animal type.
func TestSquirrelFictional(t *testing.T) {
	if v := Squirrel.Fictional(); !(v == squirrelFictional) {
		t.Fatalf("%t != %t", v, squirrelFictional)
	}
}

// TestSquirrelSpecial tests whether Squirrel is a special Animal Crossing animal type.
func TestSquirrelSpecial(t *testing.T) {
	if v := Squirrel.Special(); !(v == squirrelSpecial) {
		t.Fatalf("%t != %t", v, squirrelSpecial)
	}
}

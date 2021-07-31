package animal

import "testing"

// TestOtterId tests whether Otter has the correct ID.
func TestOtterId(t *testing.T) {
	if v := Otter.Id(); !(v == otterId) {
		t.Fatalf("%s != %s", v, otterId)
	}	
}

// TestOtterFictional test whether Otter is a fictional Animal Crossing animal type.
func TestOtterFictional(t *testing.T) {
	if v := Otter.Fictional(); !(v == otterFictional) {
		t.Fatalf("%t != %t", v, otterFictional)
	}
}

// TestOtterSpecial tests whether Otter is a special Animal Crossing animal type.
func TestOtterSpecial(t *testing.T) {
	if v := Otter.Special(); !(v == otterSpecial) {
		t.Fatalf("%t != %t", v, otterSpecial)
	}
}

package animal

import "testing"

// TestSkunkId tests whether Skunk has the correct ID.
func TestSkunkId(t *testing.T) {
	if v := Skunk.Id(); !(v == skunkId) {
		t.Fatalf("%s != %s", v, skunkId)
	}	
}

// TestSkunkFictional test whether Skunk is a fictional Animal Crossing animal type.
func TestSkunkFictional(t *testing.T) {
	if v := Skunk.Fictional(); !(v == skunkFictional) {
		t.Fatalf("%t != %t", v, skunkFictional)
	}
}

// TestSkunkSpecial tests whether Skunk is a special Animal Crossing animal type.
func TestSkunkSpecial(t *testing.T) {
	if v := Skunk.Special(); !(v == skunkSpecial) {
		t.Fatalf("%t != %t", v, skunkSpecial)
	}
}

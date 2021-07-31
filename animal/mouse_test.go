package animal

import "testing"

// TestMouseId tests whether Mouse has the correct ID.
func TestMouseId(t *testing.T) {
	if v := Mouse.Id(); !(v == mouseId) {
		t.Fatalf("%s != %s", v, mouseId)
	}	
}

// TestMouseFictional test whether Mouse is a fictional Animal Crossing animal type.
func TestMouseFictional(t *testing.T) {
	if v := Mouse.Fictional(); !(v == mouseFictional) {
		t.Fatalf("%t != %t", v, mouseFictional)
	}
}

// TestMouseSpecial tests whether Mouse is a special Animal Crossing animal type.
func TestMouseSpecial(t *testing.T) {
	if v := Mouse.Special(); !(v == mouseSpecial) {
		t.Fatalf("%t != %t", v, mouseSpecial)
	}
}

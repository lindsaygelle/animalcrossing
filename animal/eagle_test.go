package animal

import "testing"

// TestEagleId tests whether Eagle has the correct ID.
func TestEagleId(t *testing.T) {
	if v := Eagle.Id(); !(v == eagleId) {
		t.Fatalf("%s != %s", v, eagleId)
	}	
}

// TestEagleFictional test whether Eagle is a fictional Animal Crossing animal type.
func TestEagleFictional(t *testing.T) {
	if v := Eagle.Fictional(); !(v == eagleFictional) {
		t.Fatalf("%t != %t", v, eagleFictional)
	}
}

// TestEagleSpecial tests whether Eagle is a special Animal Crossing animal type.
func TestEagleSpecial(t *testing.T) {
	if v := Eagle.Special(); !(v == eagleSpecial) {
		t.Fatalf("%t != %t", v, eagleSpecial)
	}
}

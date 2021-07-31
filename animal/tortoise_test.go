package animal

import "testing"

// TestTortoiseId tests whether Tortoise has the correct ID.
func TestTortoiseId(t *testing.T) {
	if v := Tortoise.Id(); !(v == tortoiseId) {
		t.Fatalf("%s != %s", v, tortoiseId)
	}	
}

// TestTortoiseFictional test whether Tortoise is a fictional Animal Crossing animal type.
func TestTortoiseFictional(t *testing.T) {
	if v := Tortoise.Fictional(); !(v == tortoiseFictional) {
		t.Fatalf("%t != %t", v, tortoiseFictional)
	}
}

// TestTortoiseSpecial tests whether Tortoise is a special Animal Crossing animal type.
func TestTortoiseSpecial(t *testing.T) {
	if v := Tortoise.Special(); !(v == tortoiseSpecial) {
		t.Fatalf("%t != %t", v, tortoiseSpecial)
	}
}

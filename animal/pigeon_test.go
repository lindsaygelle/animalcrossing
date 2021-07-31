package animal

import "testing"

// TestPigeonId tests whether Pigeon has the correct ID.
func TestPigeonId(t *testing.T) {
	if v := Pigeon.Id(); !(v == pigeonId) {
		t.Fatalf("%s != %s", v, pigeonId)
	}	
}

// TestPigeonFictional test whether Pigeon is a fictional Animal Crossing animal type.
func TestPigeonFictional(t *testing.T) {
	if v := Pigeon.Fictional(); !(v == pigeonFictional) {
		t.Fatalf("%t != %t", v, pigeonFictional)
	}
}

// TestPigeonSpecial tests whether Pigeon is a special Animal Crossing animal type.
func TestPigeonSpecial(t *testing.T) {
	if v := Pigeon.Special(); !(v == pigeonSpecial) {
		t.Fatalf("%t != %t", v, pigeonSpecial)
	}
}

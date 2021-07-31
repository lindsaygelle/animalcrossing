package animal

import "testing"

// TestSnowId tests whether Snow has the correct ID.
func TestSnowId(t *testing.T) {
	if v := Snow.Id(); !(v == snowId) {
		t.Fatalf("%s != %s", v, snowId)
	}	
}

// TestSnowFictional test whether Snow is a fictional Animal Crossing animal type.
func TestSnowFictional(t *testing.T) {
	if v := Snow.Fictional(); !(v == snowFictional) {
		t.Fatalf("%t != %t", v, snowFictional)
	}
}

// TestSnowSpecial tests whether Snow is a special Animal Crossing animal type.
func TestSnowSpecial(t *testing.T) {
	if v := Snow.Special(); !(v == snowSpecial) {
		t.Fatalf("%t != %t", v, snowSpecial)
	}
}

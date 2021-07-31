package animal

import "testing"

// TestReindeerId tests whether Reindeer has the correct ID.
func TestReindeerId(t *testing.T) {
	if v := Reindeer.Id(); !(v == reindeerId) {
		t.Fatalf("%s != %s", v, reindeerId)
	}	
}

// TestReindeerFictional test whether Reindeer is a fictional Animal Crossing animal type.
func TestReindeerFictional(t *testing.T) {
	if v := Reindeer.Fictional(); !(v == reindeerFictional) {
		t.Fatalf("%t != %t", v, reindeerFictional)
	}
}

// TestReindeerSpecial tests whether Reindeer is a special Animal Crossing animal type.
func TestReindeerSpecial(t *testing.T) {
	if v := Reindeer.Special(); !(v == reindeerSpecial) {
		t.Fatalf("%t != %t", v, reindeerSpecial)
	}
}

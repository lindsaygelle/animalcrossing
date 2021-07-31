package animal

import "testing"

// TestKangarooId tests whether Kangaroo has the correct ID.
func TestKangarooId(t *testing.T) {
	if v := Kangaroo.Id(); !(v == kangarooId) {
		t.Fatalf("%s != %s", v, kangarooId)
	}	
}

// TestKangarooFictional test whether Kangaroo is a fictional Animal Crossing animal type.
func TestKangarooFictional(t *testing.T) {
	if v := Kangaroo.Fictional(); !(v == kangarooFictional) {
		t.Fatalf("%t != %t", v, kangarooFictional)
	}
}

// TestKangarooSpecial tests whether Kangaroo is a special Animal Crossing animal type.
func TestKangarooSpecial(t *testing.T) {
	if v := Kangaroo.Special(); !(v == kangarooSpecial) {
		t.Fatalf("%t != %t", v, kangarooSpecial)
	}
}

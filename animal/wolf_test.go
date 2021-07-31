package animal

import "testing"

// TestWolfId tests whether Wolf has the correct ID.
func TestWolfId(t *testing.T) {
	if v := Wolf.Id(); !(v == wolfId) {
		t.Fatalf("%s != %s", v, wolfId)
	}	
}

// TestWolfFictional test whether Wolf is a fictional Animal Crossing animal type.
func TestWolfFictional(t *testing.T) {
	if v := Wolf.Fictional(); !(v == wolfFictional) {
		t.Fatalf("%t != %t", v, wolfFictional)
	}
}

// TestWolfSpecial tests whether Wolf is a special Animal Crossing animal type.
func TestWolfSpecial(t *testing.T) {
	if v := Wolf.Special(); !(v == wolfSpecial) {
		t.Fatalf("%t != %t", v, wolfSpecial)
	}
}

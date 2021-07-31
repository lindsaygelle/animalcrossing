package animal

import "testing"

// TestOstrichId tests whether Ostrich has the correct ID.
func TestOstrichId(t *testing.T) {
	if v := Ostrich.Id(); !(v == ostrichId) {
		t.Fatalf("%s != %s", v, ostrichId)
	}	
}

// TestOstrichFictional test whether Ostrich is a fictional Animal Crossing animal type.
func TestOstrichFictional(t *testing.T) {
	if v := Ostrich.Fictional(); !(v == ostrichFictional) {
		t.Fatalf("%t != %t", v, ostrichFictional)
	}
}

// TestOstrichSpecial tests whether Ostrich is a special Animal Crossing animal type.
func TestOstrichSpecial(t *testing.T) {
	if v := Ostrich.Special(); !(v == ostrichSpecial) {
		t.Fatalf("%t != %t", v, ostrichSpecial)
	}
}

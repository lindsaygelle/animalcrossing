package animal

import "testing"

// TestTigerId tests whether Tiger has the correct ID.
func TestTigerId(t *testing.T) {
	if v := Tiger.Id(); !(v == tigerId) {
		t.Fatalf("%s != %s", v, tigerId)
	}	
}

// TestTigerFictional test whether Tiger is a fictional Animal Crossing animal type.
func TestTigerFictional(t *testing.T) {
	if v := Tiger.Fictional(); !(v == tigerFictional) {
		t.Fatalf("%t != %t", v, tigerFictional)
	}
}

// TestTigerSpecial tests whether Tiger is a special Animal Crossing animal type.
func TestTigerSpecial(t *testing.T) {
	if v := Tiger.Special(); !(v == tigerSpecial) {
		t.Fatalf("%t != %t", v, tigerSpecial)
	}
}

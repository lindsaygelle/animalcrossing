package animal

import "testing"

// TestOwlId tests whether Owl has the correct ID.
func TestOwlId(t *testing.T) {
	if v := Owl.Id(); !(v == owlId) {
		t.Fatalf("%s != %s", v, owlId)
	}	
}

// TestOwlFictional test whether Owl is a fictional Animal Crossing animal type.
func TestOwlFictional(t *testing.T) {
	if v := Owl.Fictional(); !(v == owlFictional) {
		t.Fatalf("%t != %t", v, owlFictional)
	}
}

// TestOwlSpecial tests whether Owl is a special Animal Crossing animal type.
func TestOwlSpecial(t *testing.T) {
	if v := Owl.Special(); !(v == owlSpecial) {
		t.Fatalf("%t != %t", v, owlSpecial)
	}
}

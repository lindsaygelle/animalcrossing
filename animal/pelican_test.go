package animal

import "testing"

// TestPelicanId tests whether Pelican has the correct ID.
func TestPelicanId(t *testing.T) {
	if v := Pelican.Id(); !(v == pelicanId) {
		t.Fatalf("%s != %s", v, pelicanId)
	}	
}

// TestPelicanFictional test whether Pelican is a fictional Animal Crossing animal type.
func TestPelicanFictional(t *testing.T) {
	if v := Pelican.Fictional(); !(v == pelicanFictional) {
		t.Fatalf("%t != %t", v, pelicanFictional)
	}
}

// TestPelicanSpecial tests whether Pelican is a special Animal Crossing animal type.
func TestPelicanSpecial(t *testing.T) {
	if v := Pelican.Special(); !(v == pelicanSpecial) {
		t.Fatalf("%t != %t", v, pelicanSpecial)
	}
}

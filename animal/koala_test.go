package animal

import "testing"

// TestKoalaId tests whether Koala has the correct ID.
func TestKoalaId(t *testing.T) {
	if v := Koala.Id(); !(v == koalaId) {
		t.Fatalf("%s != %s", v, koalaId)
	}	
}

// TestKoalaFictional test whether Koala is a fictional Animal Crossing animal type.
func TestKoalaFictional(t *testing.T) {
	if v := Koala.Fictional(); !(v == koalaFictional) {
		t.Fatalf("%t != %t", v, koalaFictional)
	}
}

// TestKoalaSpecial tests whether Koala is a special Animal Crossing animal type.
func TestKoalaSpecial(t *testing.T) {
	if v := Koala.Special(); !(v == koalaSpecial) {
		t.Fatalf("%t != %t", v, koalaSpecial)
	}
}

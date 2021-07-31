package animal

import "testing"

// TestAlpacaId tests whether Alpaca has the correct ID.
func TestAlpacaId(t *testing.T) {
	if v := Alpaca.Id(); !(v == alpacaId) {
		t.Fatalf("%s != %s", v, alpacaId)
	}	
}

// TestAlpacaFictional test whether Alpaca is a fictional Animal Crossing animal type.
func TestAlpacaFictional(t *testing.T) {
	if v := Alpaca.Fictional(); !(v == alpacaFictional) {
		t.Fatalf("%t != %t", v, alpacaFictional)
	}
}

// TestAlpacaSpecial tests whether Alpaca is a special Animal Crossing animal type.
func TestAlpacaSpecial(t *testing.T) {
	if v := Alpaca.Special(); !(v == alpacaSpecial) {
		t.Fatalf("%t != %t", v, alpacaSpecial)
	}
}

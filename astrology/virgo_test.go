package astrology

import "testing"

// TestVirgoId tests whether Virgo has the correct ID.
func TestVirgoId(t *testing.T) {
	if v := Virgo.Id(); !(v == virgoId) {
		t.Fatalf("%s != %s", v, virgoId)
	}
}

// TestVirgoLongitude test whether Virgo has the correct longitude.
func TestVirgoLongitude(t *testing.T) {
	if v := Virgo.Longtitude(); !(v == virgoLongitude) {
		t.Fatalf("%d != %d", v, virgoLongitude)
	}
}

// TestVirgoSymbol test whether Virgo has the correct symbol.
func TestVirgoSymbol(t *testing.T) {
	if v := Virgo.Symbol(); !(v == virgoSymbol) {
		t.Fatalf("%s != %s", v, virgoSymbol)
	}
}

// TestVirgoUnicode tests whether Virgo has the correct unicode character.
func TestVirgoUnicode(t *testing.T) {
	if v, ok := Virgo.Unicode(); !(v == virgoUnicode && ok) {
		t.Fatalf("%s != %s", v, virgoUnicode)
	}
}

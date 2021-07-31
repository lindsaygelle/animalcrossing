package astrology

import "testing"

// TestAriesId tests whether Aries has the correct ID.
func TestAriesId(t *testing.T) {
	if v := Aries.Id(); !(v == ariesId) {
		t.Fatalf("%s != %s", v, ariesId)
	}
}

// TestAriesLongitude test whether Aries has the correct longitude.
func TestAriesLongitude(t *testing.T) {
	if v := Aries.Longtitude(); !(v == ariesLongitude) {
		t.Fatalf("%d != %d", v, ariesLongitude)
	}
}

// TestAriesSymbol test whether Aries has the correct symbol.
func TestAriesSymbol(t *testing.T) {
	if v := Aries.Symbol(); !(v == ariesSymbol) {
		t.Fatalf("%s != %s", v, ariesSymbol)
	}
}

// TestAriesUnicode tests whether Aries has the correct unicode character.
func TestAriesUnicode(t *testing.T) {
	if v, ok := Aries.Unicode(); !(v == ariesUnicode && ok) {
		t.Fatalf("%s != %s", v, ariesUnicode)
	}
}

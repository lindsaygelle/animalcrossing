package astrology

import "testing"

// TestLibraId tests whether Libra has the correct ID.
func TestLibraId(t *testing.T) {
	if v := Libra.Id(); !(v == libraId) {
		t.Fatalf("%s != %s", v, libraId)
	}
}

// TestLibraLongitude test whether Libra has the correct longitude.
func TestLibraLongitude(t *testing.T) {
	if v := Libra.Longtitude(); !(v == libraLongitude) {
		t.Fatalf("%d != %d", v, libraLongitude)
	}
}

// TestLibraSymbol test whether Libra has the correct symbol.
func TestLibraSymbol(t *testing.T) {
	if v := Libra.Symbol(); !(v == libraSymbol) {
		t.Fatalf("%s != %s", v, libraSymbol)
	}
}

// TestLibraUnicode tests whether Libra has the correct unicode character.
func TestLibraUnicode(t *testing.T) {
	if v, ok := Libra.Unicode(); !(v == libraUnicode && ok) {
		t.Fatalf("%s != %s", v, libraUnicode)
	}
}

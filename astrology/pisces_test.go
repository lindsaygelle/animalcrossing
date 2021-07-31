package astrology

import "testing"

// TestPiscesId tests whether Pisces has the correct ID.
func TestPiscesId(t *testing.T) {
	if v := Pisces.Id(); !(v == piscesId) {
		t.Fatalf("%s != %s", v, piscesId)
	}
}

// TestPiscesLongitude test whether Pisces has the correct longitude.
func TestPiscesLongitude(t *testing.T) {
	if v := Pisces.Longtitude(); !(v == piscesLongitude) {
		t.Fatalf("%d != %d", v, piscesLongitude)
	}
}

// TestPiscesSymbol test whether Pisces has the correct symbol.
func TestPiscesSymbol(t *testing.T) {
	if v := Pisces.Symbol(); !(v == piscesSymbol) {
		t.Fatalf("%s != %s", v, piscesSymbol)
	}
}

// TestPiscesUnicode tests whether Pisces has the correct unicode character.
func TestPiscesUnicode(t *testing.T) {
	if v, ok := Pisces.Unicode(); !(v == piscesUnicode && ok) {
		t.Fatalf("%s != %s", v, piscesUnicode)
	}
}

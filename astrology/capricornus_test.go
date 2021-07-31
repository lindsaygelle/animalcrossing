package astrology

import "testing"

// TestCapricornusId tests whether Capricornus has the correct ID.
func TestCapricornusId(t *testing.T) {
	if v := Capricornus.Id(); !(v == capricornusId) {
		t.Fatalf("%s != %s", v, capricornusId)
	}
}

// TestCapricornusLongitude test whether Capricornus has the correct longitude.
func TestCapricornusLongitude(t *testing.T) {
	if v := Capricornus.Longtitude(); !(v == capricornusLongitude) {
		t.Fatalf("%d != %d", v, capricornusLongitude)
	}
}

// TestCapricornusSymbol test whether Capricornus has the correct symbol.
func TestCapricornusSymbol(t *testing.T) {
	if v := Capricornus.Symbol(); !(v == capricornusSymbol) {
		t.Fatalf("%s != %s", v, capricornusSymbol)
	}
}

// TestCapricornusUnicode tests whether Capricornus has the correct unicode character.
func TestCapricornusUnicode(t *testing.T) {
	if v, ok := Capricornus.Unicode(); !(v == capricornusUnicode && ok) {
		t.Fatalf("%s != %s", v, capricornusUnicode)
	}
}

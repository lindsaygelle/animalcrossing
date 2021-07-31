package astrology

import "testing"

// TestTaurusId tests whether Taurus has the correct ID.
func TestTaurusId(t *testing.T) {
	if v := Taurus.Id(); !(v == taurusId) {
		t.Fatalf("%s != %s", v, taurusId)
	}
}

// TestTaurusLongitude test whether Taurus has the correct longitude.
func TestTaurusLongitude(t *testing.T) {
	if v := Taurus.Longtitude(); !(v == taurusLongitude) {
		t.Fatalf("%d != %d", v, taurusLongitude)
	}
}

// TestTaurusSymbol test whether Taurus has the correct symbol.
func TestTaurusSymbol(t *testing.T) {
	if v := Taurus.Symbol(); !(v == taurusSymbol) {
		t.Fatalf("%s != %s", v, taurusSymbol)
	}
}

// TestTaurusUnicode tests whether Taurus has the correct unicode character.
func TestTaurusUnicode(t *testing.T) {
	if v, ok := Taurus.Unicode(); !(v == taurusUnicode && ok) {
		t.Fatalf("%s != %s", v, taurusUnicode)
	}
}

package astrology

import "testing"

// TestScorpiusId tests whether Scorpius has the correct ID.
func TestScorpiusId(t *testing.T) {
	if v := Scorpius.Id(); !(v == scorpiusId) {
		t.Fatalf("%s != %s", v, scorpiusId)
	}
}

// TestScorpiusLongitude test whether Scorpius has the correct longitude.
func TestScorpiusLongitude(t *testing.T) {
	if v := Scorpius.Longtitude(); !(v == scorpiusLongitude) {
		t.Fatalf("%d != %d", v, scorpiusLongitude)
	}
}

// TestScorpiusSymbol test whether Scorpius has the correct symbol.
func TestScorpiusSymbol(t *testing.T) {
	if v := Scorpius.Symbol(); !(v == scorpiusSymbol) {
		t.Fatalf("%s != %s", v, scorpiusSymbol)
	}
}

// TestScorpiusUnicode tests whether Scorpius has the correct unicode character.
func TestScorpiusUnicode(t *testing.T) {
	if v, ok := Scorpius.Unicode(); !(v == scorpiusUnicode && ok) {
		t.Fatalf("%s != %s", v, scorpiusUnicode)
	}
}

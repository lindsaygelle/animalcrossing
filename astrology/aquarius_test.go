package astrology

import "testing"

// TestAquariusId tests whether Aquarius has the correct ID.
func TestAquariusId(t *testing.T) {
	if v := Aquarius.Id(); !(v == aquariusId) {
		t.Fatalf("%s != %s", v, aquariusId)
	}
}

// TestAquariusLongitude test whether Aquarius has the correct longitude.
func TestAquariusLongitude(t *testing.T) {
	if v := Aquarius.Longtitude(); !(v == aquariusLongitude) {
		t.Fatalf("%d != %d", v, aquariusLongitude)
	}
}

// TestAquariusSymbol test whether Aquarius has the correct symbol.
func TestAquariusSymbol(t *testing.T) {
	if v := Aquarius.Symbol(); !(v == aquariusSymbol) {
		t.Fatalf("%s != %s", v, aquariusSymbol)
	}
}

// TestAquariusUnicode tests whether Aquarius has the correct unicode character.
func TestAquariusUnicode(t *testing.T) {
	if v, ok := Aquarius.Unicode(); !(v == aquariusUnicode && ok) {
		t.Fatalf("%s != %s", v, aquariusUnicode)
	}
}

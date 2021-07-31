package astrology

import "testing"

// TestLeoId tests whether Leo has the correct ID.
func TestLeoId(t *testing.T) {
	if v := Leo.Id(); !(v == leoId) {
		t.Fatalf("%s != %s", v, leoId)
	}
}

// TestLeoLongitude test whether Leo has the correct longitude.
func TestLeoLongitude(t *testing.T) {
	if v := Leo.Longtitude(); !(v == leoLongitude) {
		t.Fatalf("%d != %d", v, leoLongitude)
	}
}

// TestLeoSymbol test whether Leo has the correct symbol.
func TestLeoSymbol(t *testing.T) {
	if v := Leo.Symbol(); !(v == leoSymbol) {
		t.Fatalf("%s != %s", v, leoSymbol)
	}
}

// TestLeoUnicode tests whether Leo has the correct unicode character.
func TestLeoUnicode(t *testing.T) {
	if v, ok := Leo.Unicode(); !(v == leoUnicode && ok) {
		t.Fatalf("%s != %s", v, leoUnicode)
	}
}

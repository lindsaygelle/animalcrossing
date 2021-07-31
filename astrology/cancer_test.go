package astrology

import "testing"

// TestCancerId tests whether Cancer has the correct ID.
func TestCancerId(t *testing.T) {
	if v := Cancer.Id(); !(v == cancerId) {
		t.Fatalf("%s != %s", v, cancerId)
	}
}

// TestCancerLongitude test whether Cancer has the correct longitude.
func TestCancerLongitude(t *testing.T) {
	if v := Cancer.Longtitude(); !(v == cancerLongitude) {
		t.Fatalf("%d != %d", v, cancerLongitude)
	}
}

// TestCancerSymbol test whether Cancer has the correct symbol.
func TestCancerSymbol(t *testing.T) {
	if v := Cancer.Symbol(); !(v == cancerSymbol) {
		t.Fatalf("%s != %s", v, cancerSymbol)
	}
}

// TestCancerUnicode tests whether Cancer has the correct unicode character.
func TestCancerUnicode(t *testing.T) {
	if v, ok := Cancer.Unicode(); !(v == cancerUnicode && ok) {
		t.Fatalf("%s != %s", v, cancerUnicode)
	}
}

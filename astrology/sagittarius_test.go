package astrology

import "testing"

// TestSagittariusId tests whether Sagittarius has the correct ID.
func TestSagittariusId(t *testing.T) {
	if v := Sagittarius.Id(); !(v == sagittariusId) {
		t.Fatalf("%s != %s", v, sagittariusId)
	}
}

// TestSagittariusLongitude test whether Sagittarius has the correct longitude.
func TestSagittariusLongitude(t *testing.T) {
	if v := Sagittarius.Longtitude(); !(v == sagittariusLongitude) {
		t.Fatalf("%d != %d", v, sagittariusLongitude)
	}
}

// TestSagittariusSymbol test whether Sagittarius has the correct symbol.
func TestSagittariusSymbol(t *testing.T) {
	if v := Sagittarius.Symbol(); !(v == sagittariusSymbol) {
		t.Fatalf("%s != %s", v, sagittariusSymbol)
	}
}

// TestSagittariusUnicode tests whether Sagittarius has the correct unicode character.
func TestSagittariusUnicode(t *testing.T) {
	if v, ok := Sagittarius.Unicode(); !(v == sagittariusUnicode && ok) {
		t.Fatalf("%s != %s", v, sagittariusUnicode)
	}
}

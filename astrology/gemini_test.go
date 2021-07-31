package astrology

import "testing"

// TestGeminiId tests whether Gemini has the correct ID.
func TestGeminiId(t *testing.T) {
	if v := Gemini.Id(); !(v == geminiId) {
		t.Fatalf("%s != %s", v, geminiId)
	}
}

// TestGeminiLongitude test whether Gemini has the correct longitude.
func TestGeminiLongitude(t *testing.T) {
	if v := Gemini.Longtitude(); !(v == geminiLongitude) {
		t.Fatalf("%d != %d", v, geminiLongitude)
	}
}

// TestGeminiSymbol test whether Gemini has the correct symbol.
func TestGeminiSymbol(t *testing.T) {
	if v := Gemini.Symbol(); !(v == geminiSymbol) {
		t.Fatalf("%s != %s", v, geminiSymbol)
	}
}

// TestGeminiUnicode tests whether Gemini has the correct unicode character.
func TestGeminiUnicode(t *testing.T) {
	if v, ok := Gemini.Unicode(); !(v == geminiUnicode && ok) {
		t.Fatalf("%s != %s", v, geminiUnicode)
	}
}

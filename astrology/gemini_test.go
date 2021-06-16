package astrology

import "testing"

func TestGeminiName(t *testing.T) {
	if ok := Gemini.Name() == gemini; !ok {
		t.Fatalf("%s != %s", Gemini.Name(), gemini)
	}
}

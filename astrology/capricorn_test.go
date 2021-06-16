package astrology

import "testing"

func TestCapricornName(t *testing.T) {
	if ok := Capricorn.Name() == capricorn; !ok {
		t.Fatalf("%s != %s", Capricorn.Name(), capricorn)
	}
}

package astrology

import "testing"

func TestAriesName(t *testing.T) {
	if ok := Aries.Name() == aries; !ok {
		t.Fatalf("%s != %s", Aries.Name(), aries)
	}
}

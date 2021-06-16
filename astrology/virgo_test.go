package astrology

import "testing"

func TestVirgoName(t *testing.T) {
	if ok := Virgo.Name() == virgo; !ok {
		t.Fatalf("%s != %s", Virgo.Name(), virgo)
	}
}

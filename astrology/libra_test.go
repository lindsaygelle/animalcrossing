package astrology

import "testing"

func TestLibraName(t *testing.T) {
	if ok := Libra.Name() == libra; !ok {
		t.Fatalf("%s != %s", Libra.Name(), libra)
	}
}

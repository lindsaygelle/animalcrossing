package astrology

import "testing"

func TestPiscesName(t *testing.T) {
	if ok := Pisces.Name() == pisces; !ok {
		t.Fatalf("%s != %s", Pisces.Name(), pisces)
	}
}

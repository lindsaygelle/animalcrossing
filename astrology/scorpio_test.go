package astrology

import "testing"

func TestScorpioName(t *testing.T) {
	if ok := Scorpio.Name() == scorpio; !ok {
		t.Fatalf("%s != %s", Scorpio.Name(), scorpio)
	}
}

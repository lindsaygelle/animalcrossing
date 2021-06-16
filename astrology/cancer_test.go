package astrology

import "testing"

func TestCancerName(t *testing.T) {
	if ok := Cancer.Name() == cancer; !ok {
		t.Fatalf("%s != %s", Cancer.Name(), cancer)
	}
}

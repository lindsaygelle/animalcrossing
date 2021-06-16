package astrology

import "testing"

func TestTaurusName(t *testing.T) {
	if ok := Taurus.Name() == taurus; !ok {
		t.Fatalf("%s != %s", Taurus.Name(), taurus)
	}
}

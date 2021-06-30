package species

import "testing"

func TestCanisLupusFamiliaris(t *testing.T) {
	var s string = "Canis lupus familiaris"
	if ok := canisLupusFamiliaris == s; !ok {
		t.Fatalf("canisLupusFamiliaris != %s", s)
	}
}

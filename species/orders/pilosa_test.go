package orders

import "testing"

func TestPilosa(t *testing.T) {
	var s string = "Pilosa"
	if ok := pilosa == s; !ok {
		t.Fatalf("pilosa != %s", s)
	}
}

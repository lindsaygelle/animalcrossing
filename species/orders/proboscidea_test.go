package orders

import "testing"

func TestProboscidea(t *testing.T) {
	var s string = "Proboscidea"
	if ok := proboscidea == s; !ok {
		t.Fatalf("proboscidea != %s", s)
	}
}

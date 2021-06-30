package orders

import "testing"

func TestPrimates(t *testing.T) {
	var s string = "Primates"
	if ok := primates == s; !ok {
		t.Fatalf("primates != %s", s)
	}
}

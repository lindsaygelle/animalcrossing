package conservations

import "testing"

func TestDomesticated(t *testing.T) {
	var s string = "Domesticated"
	if ok := domesticated == s; !ok {
		t.Fatalf("domesticated != %s", s)
	}
}

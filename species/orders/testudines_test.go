package orders

import "testing"

func TestTestudines(t *testing.T) {
	var s string = "Testudines"
	if ok := testudines == s; !ok {
		t.Fatalf("testudines != %s", s)
	}
}

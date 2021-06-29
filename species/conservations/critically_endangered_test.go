package conservations

import "testing"

func TestCriticallyEndangered(t *testing.T) {
	var s string = "Critically Endangered"
	if ok := criticallyEndangered == s; !ok {
		t.Fatalf("criticallyEndangered != %s", s)
	}
}

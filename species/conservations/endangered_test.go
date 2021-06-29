package conservations

import "testing"

func TestEndangered(t *testing.T) {
	var s string = "Endangered"
	if ok := endangered == s; !ok {
		t.Fatalf("endangered != %s", s)
	}
}

package families

import "testing"

func TestBovidae(t *testing.T) {
	var s string = "Bovidae"
	if ok := bovidae == s; !ok {
		t.Fatalf("bovidae != %s", s)
	}
}

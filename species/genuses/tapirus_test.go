package genuses

import "testing"

func TestTapirus(t *testing.T) {
	var s string = "Tapirus"
	if ok := tapirus == s; !ok {
		t.Fatalf("tapirus != %s", s)
	}
}

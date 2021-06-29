package genuses

import "testing"

func TestGenusTapirus(t *testing.T) {
	var s string = "Tapirus"
	if ok := tapirus == s; !ok {
		t.Fatalf("tapirus != %s", s)
	}
}

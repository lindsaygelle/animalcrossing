package conservations

import "testing"

func TestUnknown(t *testing.T) {
	var s string = "Unknown"
	if ok := unknown == s; !ok {
		t.Fatalf("unknown != %s", s)
	}
}

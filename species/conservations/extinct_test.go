package conservations

import "testing"

func TestExtinct(t *testing.T) {
	var s string = "Extinct"
	if ok := extinct == s; !ok {
		t.Fatalf("extinct != %s", s)
	}
}

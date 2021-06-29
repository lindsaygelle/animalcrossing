package genuses

import "testing"

func TestGenusPanthera(t *testing.T) {
	var s string = "Panthera"
	if ok := panthera == s; !ok {
		t.Fatalf("panthera != %s", s)
	}
}

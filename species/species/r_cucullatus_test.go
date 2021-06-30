package species

import "testing"

func TestRCucullatus(t *testing.T) {
	var s string = "R. cucullatus"
	if ok := rCucullatus == s; !ok {
		t.Fatalf("rCucullatus != %s", s)
	}
}

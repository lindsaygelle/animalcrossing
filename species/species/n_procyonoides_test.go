package species

import "testing"

func TestNProcyonoides(t *testing.T) {
	var s string = "N. procyonoides"
	if ok := nProcyonoides == s; !ok {
		t.Fatalf("nProcyonoides != %s", s)
	}
}

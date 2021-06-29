package families

import "testing"

func TestMephitidae(t *testing.T) {
	var s string = "Mephitidae"
	if ok := mephitidae == s; !ok {
		t.Fatalf("mephitidae != %s", s)
	}
}

package families

import "testing"

func TestMacropodidae(t *testing.T) {
	var s string = "Macropodidae"
	if ok := macropodidae == s; !ok {
		t.Fatalf("macropodidae != %s", s)
	}
}

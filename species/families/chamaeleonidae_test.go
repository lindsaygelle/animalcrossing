package families

import "testing"

func TestFamilyChamaeleonidae(t *testing.T) {
	var s string = "Chamaeleonidae"
	if ok := chamaeleonidae == s; !ok {
		t.Fatalf("chamaeleonidae != %s", s)
	}
}

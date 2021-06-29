package families

import "testing"

func TestColumbidae(t *testing.T) {
	var s string = "Columbidae"
	if ok := columbidae == s; !ok {
		t.Fatalf("columbidae != %s", s)
	}
}

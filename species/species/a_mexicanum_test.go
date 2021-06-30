package species

import "testing"

func TestAMexicanum(t *testing.T) {
	var s string = "A. mexicanum"
	if ok := aMexicanum == s; !ok {
		t.Fatalf("aMexicanum != %s", s)
	}
}

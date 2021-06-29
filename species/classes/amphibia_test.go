package classes

import "testing"

func TestAmphibia(t *testing.T) {
	var s string = "Amphibia"
	if ok := amphibia == s; !ok {
		t.Fatalf("amphibia != %s", s)
	}
}

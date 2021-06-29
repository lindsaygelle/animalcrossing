package families

import "testing"

func TestCamelidae(t *testing.T) {
	var s string = "Camelidae"
	if ok := camelidae == s; !ok {
		t.Fatalf("il != %s", s)
	}
}

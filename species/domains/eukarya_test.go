package domains

import "testing"

func TestDomainEukarya(t *testing.T) {
	var s string = "Eukarya"
	if ok := eukarya == s; !ok {
		t.Fatalf("eukarya != %s", s)
	}
}

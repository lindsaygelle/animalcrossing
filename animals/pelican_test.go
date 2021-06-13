package animals

import "testing"

func TestPelicanName(t *testing.T) {
	if ok := Pelican.Name() == pelican; !ok {
		t.Fatal("Pelican.Name() != pelican")
	}
}

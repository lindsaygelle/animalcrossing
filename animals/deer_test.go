package animals

import "testing"

func TestDeerName(t *testing.T) {
	if ok := Deer.Name() == deer; !ok {
		t.Fatal("Deer.Name() != deer")
	}
}

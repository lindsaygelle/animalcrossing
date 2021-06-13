package animals

import "testing"

func TestReptileName(t *testing.T) {
	if ok := Reptile.Name() == reptile; !ok {
		t.Fatal("Reptile.Name() != reptile")
	}
}

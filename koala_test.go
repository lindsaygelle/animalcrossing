package animalcrossing

import "testing"

func TestKoalaName(t *testing.T) {
	if ok := Koala.Name() == koala; !ok {
		t.Fatal("Koala.Name() != koala")
	}
}

package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestEugeneAnimal(t *testing.T) {
	var s string = animals.Koala.Name()
	if ok := Eugene.Animal() == s; !ok {
		t.Fatalf("%s != %s", Eugene.Animal(), s)
	}
}

func TestEugeneName(t *testing.T) {
	if ok := Eugene.Name() == eugene; !ok {
		t.Fatalf("%s != %s", Eugene.Name(), eugene)
	}
}

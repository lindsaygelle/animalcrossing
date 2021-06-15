package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBlancaAnimal(t *testing.T) {
	var s string = animals.Cat.Name()
	if ok := Blanca.Animal() == s; !ok {
		t.Fatalf("%s != %s", Blanca.Animal(), s)
	}
}

func TestBlancaName(t *testing.T) {
	if ok := Blanca.Name() == blanca; !ok {
		t.Fatalf("%s != %s", Blanca.Name(), blanca)
	}
}

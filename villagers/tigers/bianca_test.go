package tigers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBiancaAnimal(t *testing.T) {
	var s string = animals.Tiger.Name()
	if ok := Bianca.Animal() == s; !ok {
		t.Fatalf("%s != %s", Bianca.Animal(), s)
	}
}

func TestBiancaName(t *testing.T) {
	if ok := Bianca.Name() == bianca; !ok {
		t.Fatalf("%s != %s", Bianca.Name(), bianca)
	}
}

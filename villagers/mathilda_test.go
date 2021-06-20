package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestMathildaAnimal(t *testing.T) {
	var s string = animals.Kangaroo.Name()
	if ok := Mathilda.Animal() == s; !ok {
		t.Fatalf("%s != %s", Mathilda.Animal(), s)
	}
}

func TestMathildaName(t *testing.T) {
	if ok := Mathilda.Name() == mathilda; !ok {
		t.Fatalf("%s != %s", Mathilda.Name(), mathilda)
	}
}

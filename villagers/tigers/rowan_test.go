package tigers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestRowanAnimal(t *testing.T) {
	var s string = animals.Tiger.Name()
	if ok := Rowan.Animal() == s; !ok {
		t.Fatalf("%s != %s", Rowan.Animal(), s)
	}
}

func TestRowanName(t *testing.T) {
	if ok := Rowan.Name() == rowan; !ok {
		t.Fatalf("%s != %s", Rowan.Name(), rowan)
	}
}

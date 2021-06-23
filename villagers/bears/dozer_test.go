package bears

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestDozerAnimal(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := Dozer.Animal() == s; !ok {
		t.Fatalf("%s != %s", Dozer.Animal(), s)
	}
}

func TestDozerName(t *testing.T) {
	if ok := Dozer.Name() == dozer; !ok {
		t.Fatalf("%s != %s", Dozer.Name(), dozer)
	}
}

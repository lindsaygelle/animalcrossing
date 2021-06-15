package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBeardoAnimal(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := Beardo.Animal() == s; !ok {
		t.Fatalf("%s != %s", Beardo.Animal(), s)
	}
}

func TestBeardoName(t *testing.T) {
	if ok := Beardo.Name() == beardo; !ok {
		t.Fatalf("%s != %s", Beardo.Name(), beardo)
	}
}

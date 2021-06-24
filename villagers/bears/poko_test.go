package bears

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPokoAnimal(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := Poko.Animal() == s; !ok {
		t.Fatalf("%s != %s", Poko.Animal(), s)
	}
}

func TestPokoName(t *testing.T) {
	if ok := Poko.Name() == poko; !ok {
		t.Fatalf("%s != %s", Poko.Name(), poko)
	}
}
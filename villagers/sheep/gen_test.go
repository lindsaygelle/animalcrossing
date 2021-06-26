package sheep

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestGenAnimal(t *testing.T) {
	var s string = animals.Sheep.Name()
	if ok := Gen.Animal() == s; !ok {
		t.Fatalf("%s != %s", Gen.Animal(), s)
	}
}

func TestGenName(t *testing.T) {
	if ok := Gen.Name() == gen; !ok {
		t.Fatalf("%s != %s", Gen.Name(), gen)
	}
}

package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestRaymondAnimal(t *testing.T) {
	var s string = animals.Cat.Name()
	if ok := Raymond.Animal() == s; !ok {
		t.Fatalf("%s != %s", Raymond.Animal(), s)
	}
}

func TestRaymondName(t *testing.T) {
	if ok := Raymond.Name() == raymond; !ok {
		t.Fatalf("%s != %s", Raymond.Name(), raymond)
	}
}

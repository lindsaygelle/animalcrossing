package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestKidCatAnimal(t *testing.T) {
	var s string = animals.Cat.Name()
	if ok := KidCat.Animal() == s; !ok {
		t.Fatalf("%s != %s", KidCat.Animal(), s)
	}
}

func TestKidCatName(t *testing.T) {
	if ok := KidCat.Name() == kidCat; !ok {
		t.Fatalf("%s != %s", KidCat.Name(), kidCat)
	}
}

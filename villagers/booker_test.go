package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBookerAnimal(t *testing.T) {
	var s string = animals.Dog.Name()
	if ok := Booker.Animal() == s; !ok {
		t.Fatalf("%s != %s", Booker.Animal(), s)
	}
}

func TestBookerName(t *testing.T) {
	if ok := Booker.Name() == booker; !ok {
		t.Fatalf("%s != %s", Booker.Name(), booker)
	}
}

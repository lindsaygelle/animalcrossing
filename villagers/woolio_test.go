package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestWoolioAnimal(t *testing.T) {
	var s string = animals.Sheep.Name()
	if ok := Woolio.Animal() == s; !ok {
		t.Fatalf("%s != %s", Woolio.Animal(), s)
	}
}

func TestWoolioName(t *testing.T) {
	if ok := Woolio.Name() == woolio; !ok {
		t.Fatalf("%s != %s", Woolio.Name(), woolio)
	}
}

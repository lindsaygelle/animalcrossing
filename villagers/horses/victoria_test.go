package horses

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestVictoriaAnimal(t *testing.T) {
	var s string = animals.Horse.Name()
	if ok := Victoria.Animal() == s; !ok {
		t.Fatalf("%s != %s", Victoria.Animal(), s)
	}
}

func TestVictoriaName(t *testing.T) {
	if ok := Victoria.Name() == victoria; !ok {
		t.Fatalf("%s != %s", Victoria.Name(), victoria)
	}
}

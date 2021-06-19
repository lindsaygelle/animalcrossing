package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPapiAnimal(t *testing.T) {
	var s string = animals.Horse.Name()
	if ok := Papi.Animal() == s; !ok {
		t.Fatalf("%s != %s", Papi.Animal(), s)
	}
}

func TestPapiName(t *testing.T) {
	if ok := Papi.Name() == papi; !ok {
		t.Fatalf("%s != %s", Papi.Name(), papi)
	}
}

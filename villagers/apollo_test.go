package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestApolloAnimal(t *testing.T) {
	var s string = animals.Eagle.Name()
	if ok := Apollo.Animal() == s; !ok {
		t.Fatalf("%s != %s", Apollo.Animal(), s)
	}
}

func TestApolloName(t *testing.T) {
	if ok := Apollo.Name() == apollo; !ok {
		t.Fatalf("%s != %s", Apollo.Name(), apollo)
	}
}

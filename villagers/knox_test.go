package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestKnoxAnimal(t *testing.T) {
	var s string = animals.Chicken.Name()
	if ok := Knox.Animal() == s; !ok {
		t.Fatalf("%s != %s", Knox.Animal(), s)
	}
}

func TestKnoxName(t *testing.T) {
	if ok := Knox.Name() == knox; !ok {
		t.Fatalf("%s != %s", Knox.Name(), knox)
	}
}

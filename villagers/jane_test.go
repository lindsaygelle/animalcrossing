package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestJaneAnimal(t *testing.T) {
	var s string = animals.Gorilla.Name()
	if ok := Jane.Animal() == s; !ok {
		t.Fatalf("%s != %s", Jane.Animal(), s)
	}
}

func TestJaneName(t *testing.T) {
	if ok := Jane.Name() == jane; !ok {
		t.Fatalf("%s != %s", Jane.Name(), jane)
	}
}

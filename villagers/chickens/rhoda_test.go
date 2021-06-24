package chickens

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestRhodaAnimal(t *testing.T) {
	var s string = animals.Chicken.Name()
	if ok := Rhoda.Animal() == s; !ok {
		t.Fatalf("%s != %s", Rhoda.Animal(), s)
	}
}

func TestRhodaName(t *testing.T) {
	if ok := Rhoda.Name() == rhoda; !ok {
		t.Fatalf("%s != %s", Rhoda.Name(), rhoda)
	}
}

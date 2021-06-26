package sheep

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBaabaraAnimal(t *testing.T) {
	var s string = animals.Sheep.Name()
	if ok := Baabara.Animal() == s; !ok {
		t.Fatalf("%s != %s", Baabara.Animal(), s)
	}
}

func TestBaabaraName(t *testing.T) {
	if ok := Baabara.Name() == baabara; !ok {
		t.Fatalf("%s != %s", Baabara.Name(), baabara)
	}
}

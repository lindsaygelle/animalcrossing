package eagles

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestSterlingAnimal(t *testing.T) {
	var s string = animals.Eagle.Name()
	if ok := Sterling.Animal() == s; !ok {
		t.Fatalf("%s != %s", Sterling.Animal(), s)
	}
}

func TestSterlingName(t *testing.T) {
	if ok := Sterling.Name() == sterling; !ok {
		t.Fatalf("%s != %s", Sterling.Name(), sterling)
	}
}

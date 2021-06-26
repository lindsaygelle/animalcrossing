package horses

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestJulianAnimal(t *testing.T) {
	var s string = animals.Horse.Name()
	if ok := Julian.Animal() == s; !ok {
		t.Fatalf("%s != %s", Julian.Animal(), s)
	}
}

func TestJulianName(t *testing.T) {
	if ok := Julian.Name() == julian; !ok {
		t.Fatalf("%s != %s", Julian.Name(), julian)
	}
}

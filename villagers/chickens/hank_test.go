package chickens

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestHankAnimal(t *testing.T) {
	var s string = animals.Chicken.Name()
	if ok := Hank.Animal() == s; !ok {
		t.Fatalf("%s != %s", Hank.Animal(), s)
	}
}

func TestHankName(t *testing.T) {
	if ok := Hank.Name() == hank; !ok {
		t.Fatalf("%s != %s", Hank.Name(), hank)
	}
}

package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestTBoneAnimal(t *testing.T) {
	var s string = animals.Bull.Name()
	if ok := TBone.Animal() == s; !ok {
		t.Fatalf("%s != %s", TBone.Animal(), s)
	}
}

func TestTBoneName(t *testing.T) {
	if ok := TBone.Name() == tBone; !ok {
		t.Fatalf("%s != %s", TBone.Name(), tBone)
	}
}

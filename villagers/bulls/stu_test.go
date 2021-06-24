package bulls

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestStuAnimal(t *testing.T) {
	var s string = animals.Bull.Name()
	if ok := Stu.Animal() == s; !ok {
		t.Fatalf("%s != %s", Stu.Animal(), s)
	}
}

func TestStuName(t *testing.T) {
	if ok := Stu.Name() == stu; !ok {
		t.Fatalf("%s != %s", Stu.Name(), stu)
	}
}

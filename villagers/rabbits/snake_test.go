package rabbits

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestSnakeAnimal(t *testing.T) {
	var s string = animals.Rabbit.Name()
	if ok := Snake.Animal() == s; !ok {
		t.Fatalf("%s != %s", Snake.Animal(), s)
	}
}

func TestSnakeName(t *testing.T) {
	if ok := Snake.Name() == snake; !ok {
		t.Fatalf("%s != %s", Snake.Name(), snake)
	}
}

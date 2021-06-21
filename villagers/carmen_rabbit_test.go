package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestCarmenRabbitAnimal(t *testing.T) {
	var s string = animals.Rabbit.Name()
	if ok := CarmenRabbit.Animal() == s; !ok {
		t.Fatalf("%s != %s", CarmenRabbit.Animal(), s)
	}
}

func TestCarmenRabbitName(t *testing.T) {
	if ok := CarmenRabbit.Name() == carmen; !ok {
		t.Fatalf("%s != %s", CarmenRabbit.Name(), carmen)
	}
}

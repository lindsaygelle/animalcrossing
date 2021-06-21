package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestClaudeAnimal(t *testing.T) {
	var s string = animals.Rabbit.Name()
	if ok := Claude.Animal() == s; !ok {
		t.Fatalf("%s != %s", Claude.Animal(), s)
	}
}

func TestClaudeName(t *testing.T) {
	if ok := Claude.Name() == claude; !ok {
		t.Fatalf("%s != %s", Claude.Name(), claude)
	}
}

package mouse

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animal"
)

func TestVillagers(t *testing.T) {
	for _, villager := range Villagers {
		if ok := villager.Animal.Key == animal.Mouse; !ok {
			t.Fatalf("%s.Animal.Key != %d", villager.Id, animal.Mouse)
		}
	}
}

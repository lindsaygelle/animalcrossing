package tiger

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animal"
)

func TestVillagers(t *testing.T) {
	for _, villager := range Villagers {
		if ok := villager.Animal.Key == animal.Tiger; !ok {
			t.Fatalf("%s.Animal.Key != %d", villager.Id, animal.Tiger)
		}
	}
}

package rodney

import (
	"github.com/lindsaygelle/animalcrossing/animal/hamster"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "rodney"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Rodney.
	Villager = villager.Villager{
		Animal:  hamster.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)

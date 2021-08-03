package francine

import (
	"github.com/lindsaygelle/animalcrossing/animal/rabbit"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "francine"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Francine.
	Villager = villager.Villager{
		Animal:  rabbit.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)

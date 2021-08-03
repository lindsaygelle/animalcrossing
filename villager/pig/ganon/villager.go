package ganon

import (
	"github.com/lindsaygelle/animalcrossing/animal/pig"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "ganon"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Ganon.
	Villager = villager.Villager{
		Animal:  pig.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)

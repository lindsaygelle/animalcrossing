package paolo

import (
	"github.com/lindsaygelle/animalcrossing/animal/elephant"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "paolo"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Paolo.
	Villager = villager.Villager{
		Animal:  elephant.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)

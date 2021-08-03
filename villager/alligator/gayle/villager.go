package gayle

import (
	"github.com/lindsaygelle/animalcrossing/animal/alligator"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "gayle"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Gayle.
	Villager = villager.Villager{
		Animal:  alligator.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)

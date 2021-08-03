package marina

import (
	"github.com/lindsaygelle/animalcrossing/animal/octopus"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "marina"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Marina.
	Villager = villager.Villager{
		Animal:  octopus.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)

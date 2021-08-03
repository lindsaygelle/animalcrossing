package frobert

import (
	"github.com/lindsaygelle/animalcrossing/animal/frog"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "frobert"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Frobert.
	Villager = villager.Villager{
		Animal:  frog.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)

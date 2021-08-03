package monique

import (
	"github.com/lindsaygelle/animalcrossing/animal/cat"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "monique"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Monique.
	Villager = villager.Villager{
		Animal:  cat.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)

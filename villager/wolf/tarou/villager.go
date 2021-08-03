package tarou

import (
	"github.com/lindsaygelle/animalcrossing/animal/wolf"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "tarou"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Tarou.
	Villager = villager.Villager{
		Animal:  wolf.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)

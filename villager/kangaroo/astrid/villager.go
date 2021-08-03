package astrid

import (
	"github.com/lindsaygelle/animalcrossing/animal/kangaroo"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "astrid"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Astrid.
	Villager = villager.Villager{
		Animal:  kangaroo.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)

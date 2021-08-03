package nibbles

import (
	"github.com/lindsaygelle/animalcrossing/animal/squirrel"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "nibbles"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Nibbles.
	Villager = villager.Villager{
		Animal:  squirrel.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)

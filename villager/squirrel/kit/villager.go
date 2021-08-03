package kit

import (
	"github.com/lindsaygelle/animalcrossing/animal/squirrel"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "kit"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Kit.
	Villager = villager.Villager{
		Animal:  squirrel.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)

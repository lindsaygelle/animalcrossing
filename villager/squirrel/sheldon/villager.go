package sheldon

import (
	"github.com/lindsaygelle/animalcrossing/animal/squirrel"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "sheldon"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Sheldon.
	Villager = villager.Villager{
		Animal:  squirrel.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)

package coach

import (
	"github.com/lindsaygelle/animalcrossing/animal/bull"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "coach"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Coach.
	Villager = villager.Villager{
		Animal:  bull.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)

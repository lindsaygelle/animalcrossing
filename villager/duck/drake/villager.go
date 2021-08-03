package drake

import (
	"github.com/lindsaygelle/animalcrossing/animal/duck"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "drake"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Drake.
	Villager = villager.Villager{
		Animal:  duck.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)

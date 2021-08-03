package erik

import (
	"github.com/lindsaygelle/animalcrossing/animal/deer"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "erik"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Erik.
	Villager = villager.Villager{
		Animal:  deer.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)

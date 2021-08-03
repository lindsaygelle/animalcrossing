package june

import (
	"github.com/lindsaygelle/animalcrossing/animal/bearcub"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "june"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for June.
	Villager = villager.Villager{
		Animal:  bearcub.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)

package chai

import (
	"github.com/lindsaygelle/animalcrossing/animal/elephant"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "chai"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Chai.
	Villager = villager.Villager{
		Animal:  elephant.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)

package chai

import (
	"github.com/lindsaygelle/animalcrossing/animal/elephant"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "chai"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Chai.
	Villager = villager.Villager{
		Animal:      elephant.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Chai,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

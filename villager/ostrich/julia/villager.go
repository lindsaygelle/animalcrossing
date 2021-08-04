package julia

import (
	"github.com/lindsaygelle/animalcrossing/animal/ostrich"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "julia"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Julia.
	Villager = villager.Villager{
		Animal:      ostrich.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Julia,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

package octavian

import (
	"github.com/lindsaygelle/animalcrossing/animal/octopus"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "ocp00"
)

const (
	gender string = "male"
)

const (
	id string = "octavian"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Octavian.
	Villager = villager.Villager{
		Animal:      octopus.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Octavian,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

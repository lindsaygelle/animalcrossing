package tammi

import (
	"github.com/lindsaygelle/animalcrossing/animal/monkey"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "tammi"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Tammi.
	Villager = villager.Villager{
		Animal:      monkey.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Tammi,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

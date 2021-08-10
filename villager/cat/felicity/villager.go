package felicity

import (
	"github.com/lindsaygelle/animalcrossing/animal/cat"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "cat17"
)

const (
	gender string = "female"
)

const (
	id string = "felicity"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Felicity.
	Villager = villager.Villager{
		Animal:      cat.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Felicity,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

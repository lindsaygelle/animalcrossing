package melba

import (
	"github.com/lindsaygelle/animalcrossing/animal/koala"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "melba"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Melba.
	Villager = villager.Villager{
		Animal:      koala.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Melba,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

package renee

import (
	"github.com/lindsaygelle/animalcrossing/animal/rhinoceros"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "rhn08"
)

const (
	gender string = "female"
)

const (
	id string = "renee"
)

const (
	personality string = "sisterly"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Renee.
	Villager = villager.Villager{
		Animal:      rhinoceros.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Renee,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

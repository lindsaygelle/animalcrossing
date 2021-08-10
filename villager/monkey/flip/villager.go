package flip

import (
	"github.com/lindsaygelle/animalcrossing/animal/monkey"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "mnk06"
)

const (
	gender string = "male"
)

const (
	id string = "flip"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Flip.
	Villager = villager.Villager{
		Animal:      monkey.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Flip,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

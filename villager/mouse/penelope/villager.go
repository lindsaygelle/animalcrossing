package penelope

import (
	"github.com/lindsaygelle/animalcrossing/animal/mouse"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "mus17"
)

const (
	gender string = "female"
)

const (
	id string = "penelope"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Penelope.
	Villager = villager.Villager{
		Animal:      mouse.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Penelope,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

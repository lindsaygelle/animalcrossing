package derwin

import (
	"github.com/lindsaygelle/animalcrossing/animal/duck"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "duk08"
)

const (
	gender string = "male"
)

const (
	id string = "derwin"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Derwin.
	Villager = villager.Villager{
		Animal:      duck.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Derwin,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

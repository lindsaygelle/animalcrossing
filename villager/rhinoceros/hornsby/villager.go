package hornsby

import (
	"github.com/lindsaygelle/animalcrossing/animal/rhinoceros"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "rhn04"
)

const (
	gender string = "male"
)

const (
	id string = "hornsby"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Hornsby.
	Villager = villager.Villager{
		Animal:      rhinoceros.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Hornsby,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

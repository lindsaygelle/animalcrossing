package merengue

import (
	"github.com/lindsaygelle/animalcrossing/animal/rhinoceros"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "rhn07"
)

const (
	gender string = "female"
)

const (
	id string = "merengue"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Merengue.
	Villager = villager.Villager{
		Animal:      rhinoceros.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Merengue,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

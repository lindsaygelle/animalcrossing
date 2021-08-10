package phil

import (
	"github.com/lindsaygelle/animalcrossing/animal/ostrich"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "ost07"
)

const (
	gender string = "male"
)

const (
	id string = "phil"
)

const (
	personality string = "smug"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Phil.
	Villager = villager.Villager{
		Animal:      ostrich.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Phil,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

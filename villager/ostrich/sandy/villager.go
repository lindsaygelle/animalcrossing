package sandy

import (
	"github.com/lindsaygelle/animalcrossing/animal/ostrich"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "ost02"
)

const (
	gender string = "female"
)

const (
	id string = "sandy"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Sandy.
	Villager = villager.Villager{
		Animal:      ostrich.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Sandy,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

package ken

import (
	"github.com/lindsaygelle/animalcrossing/animal/chicken"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "chn13"
)

const (
	gender string = "male"
)

const (
	id string = "ken"
)

const (
	personality string = "smug"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Ken.
	Villager = villager.Villager{
		Animal:      chicken.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Ken,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

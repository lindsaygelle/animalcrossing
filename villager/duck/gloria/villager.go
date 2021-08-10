package gloria

import (
	"github.com/lindsaygelle/animalcrossing/animal/duck"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "duk15"
)

const (
	gender string = "female"
)

const (
	id string = "gloria"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Gloria.
	Villager = villager.Villager{
		Animal:      duck.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Gloria,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

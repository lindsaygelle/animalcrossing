package hopper

import (
	"github.com/lindsaygelle/animalcrossing/animal/penguin"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "pgn03"
)

const (
	gender string = "male"
)

const (
	id string = "hopper"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Hopper.
	Villager = villager.Villager{
		Animal:      penguin.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Hopper,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

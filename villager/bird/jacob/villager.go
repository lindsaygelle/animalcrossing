package jacob

import (
	"github.com/lindsaygelle/animalcrossing/animal/bird"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "brd11"
)

const (
	gender string = "male"
)

const (
	id string = "jacob"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Jacob.
	Villager = villager.Villager{
		Animal:      bird.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Jacob,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

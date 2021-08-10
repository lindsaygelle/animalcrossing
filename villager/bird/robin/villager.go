package robin

import (
	"github.com/lindsaygelle/animalcrossing/animal/bird"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "brd01"
)

const (
	gender string = "female"
)

const (
	id string = "robin"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Robin.
	Villager = villager.Villager{
		Animal:      bird.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Robin,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

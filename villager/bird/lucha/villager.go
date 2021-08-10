package lucha

import (
	"github.com/lindsaygelle/animalcrossing/animal/bird"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "brd15"
)

const (
	gender string = "male"
)

const (
	id string = "lucha"
)

const (
	personality string = "smug"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Lucha.
	Villager = villager.Villager{
		Animal:      bird.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Lucha,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

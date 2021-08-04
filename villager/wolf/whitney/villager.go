package whitney

import (
	"github.com/lindsaygelle/animalcrossing/animal/wolf"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = ""
)

const (
	gender string = "female"
)

const (
	id string = "whitney"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Whitney.
	Villager = villager.Villager{
		Animal:      wolf.Animal,
		Birthday:    birthday,
		Code:        code,
		Gender:      gender,
		Id:          id,
		Key:         villager.Whitney,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

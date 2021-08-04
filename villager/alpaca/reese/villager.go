package reese

import (
	"github.com/lindsaygelle/animalcrossing/animal/alpaca"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = ""
)

const (
	gender string = "female"
)

const (
	id string = "reese"
)

const (
	personality string = ""
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Reese.
	Villager = villager.Villager{
		Animal:      alpaca.Animal,
		Birthday:    birthday,
		Code:        code,
		Gender:      gender,
		Id:          id,
		Key:         villager.Reese,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

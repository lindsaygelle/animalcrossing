package cyrus

import (
	"github.com/lindsaygelle/animalcrossing/animal/alpaca"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = ""
)

const (
	gender string = "male"
)

const (
	id string = "cyrus"
)

const (
	personality string = ""
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Cyrus.
	Villager = villager.Villager{
		Animal:      alpaca.Animal,
		Birthday:    birthday,
		Code:        code,
		Gender:      gender,
		Id:          id,
		Key:         villager.Cyrus,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

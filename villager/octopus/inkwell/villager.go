package inkwell

import (
	"github.com/lindsaygelle/animalcrossing/animal/octopus"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = ""
)

const (
	gender string = "male"
)

const (
	id string = "inkwell"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Inkwell.
	Villager = villager.Villager{
		Animal:      octopus.Animal,
		Birthday:    birthday,
		Code:        code,
		Gender:      gender,
		Id:          id,
		Key:         villager.Inkwell,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

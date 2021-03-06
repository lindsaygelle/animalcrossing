package quetzal

import (
	"github.com/lindsaygelle/animalcrossing/animal/eagle"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = ""
)

const (
	gender string = "male"
)

const (
	id string = "quetzal"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Quetzal.
	Villager = villager.Villager{
		Animal:      eagle.Animal,
		Birthday:    birthday,
		Code:        code,
		Gender:      gender,
		Id:          id,
		Key:         villager.Quetzal,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

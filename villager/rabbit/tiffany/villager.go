package tiffany

import (
	"github.com/lindsaygelle/animalcrossing/animal/rabbit"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "rbt07"
)

const (
	gender string = "female"
)

const (
	id string = "tiffany"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Tiffany.
	Villager = villager.Villager{
		Animal:      rabbit.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Tiffany,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

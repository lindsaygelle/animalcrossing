package gigi

import (
	"github.com/lindsaygelle/animalcrossing/animal/frog"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "flg16"
)

const (
	gender string = "female"
)

const (
	id string = "gigi"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Gigi.
	Villager = villager.Villager{
		Animal:      frog.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Gigi,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

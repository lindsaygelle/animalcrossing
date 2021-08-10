package bianca

import (
	"github.com/lindsaygelle/animalcrossing/animal/tiger"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "tig06"
)

const (
	gender string = "female"
)

const (
	id string = "bianca"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Bianca.
	Villager = villager.Villager{
		Animal:      tiger.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Bianca,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

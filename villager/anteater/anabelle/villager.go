package anabelle

import (
	"github.com/lindsaygelle/animalcrossing/animal/anteater"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "ant03"
)

const (
	gender string = "female"
)

const (
	id string = "anabelle"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Anabelle.
	Villager = villager.Villager{
		Animal:      anteater.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Anabelle,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

package drago

import (
	"github.com/lindsaygelle/animalcrossing/animal/alligator"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "crd08"
)

const (
	gender string = "male"
)

const (
	id string = "drago"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Drago.
	Villager = villager.Villager{
		Animal:      alligator.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Drago,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

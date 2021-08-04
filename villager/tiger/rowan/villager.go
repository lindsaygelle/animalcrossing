package rowan

import (
	"github.com/lindsaygelle/animalcrossing/animal/tiger"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "rowan"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Rowan.
	Villager = villager.Villager{
		Animal:      tiger.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Rowan,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

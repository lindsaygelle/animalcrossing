package carmen

import (
	"github.com/lindsaygelle/animalcrossing/animal/rabbit"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "carmen"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Carmen.
	Villager = villager.Villager{
		Animal:      rabbit.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Carmen,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

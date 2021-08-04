package biskit

import (
	"github.com/lindsaygelle/animalcrossing/animal/dog"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "biskit"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Biskit.
	Villager = villager.Villager{
		Animal:      dog.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Biskit,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

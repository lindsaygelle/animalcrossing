package megumi

import (
	"github.com/lindsaygelle/animalcrossing/animal/dog"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "megumi"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Megumi.
	Villager = villager.Villager{
		Animal:      dog.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Megumi,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

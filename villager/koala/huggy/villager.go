package huggy

import (
	"github.com/lindsaygelle/animalcrossing/animal/koala"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "huggy"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Huggy.
	Villager = villager.Villager{
		Animal:      koala.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Huggy,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

package marty

import (
	"github.com/lindsaygelle/animalcrossing/animal/bearcub"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "marty"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Marty.
	Villager = villager.Villager{
		Animal:      bearcub.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Marty,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

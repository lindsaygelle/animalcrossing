package mac

import (
	"github.com/lindsaygelle/animalcrossing/animal/dog"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "mac"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Mac.
	Villager = villager.Villager{
		Animal:      dog.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Mac,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

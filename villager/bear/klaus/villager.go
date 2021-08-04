package klaus

import (
	"github.com/lindsaygelle/animalcrossing/animal/bear"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "klaus"
)

const (
	personality string = "smug"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Klaus.
	Villager = villager.Villager{
		Animal:      bear.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Klaus,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

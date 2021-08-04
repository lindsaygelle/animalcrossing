package beardo

import (
	"github.com/lindsaygelle/animalcrossing/animal/bear"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "beardo"
)

const (
	personality string = "smug"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Beardo.
	Villager = villager.Villager{
		Animal:      bear.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Beardo,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

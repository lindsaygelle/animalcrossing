package rodney

import (
	"github.com/lindsaygelle/animalcrossing/animal/hamster"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "ham03"
)

const (
	gender string = "male"
)

const (
	id string = "rodney"
)

const (
	personality string = "smug"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Rodney.
	Villager = villager.Villager{
		Animal:      hamster.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Rodney,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

package chuck

import (
	"github.com/lindsaygelle/animalcrossing/animal/bull"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "chuck"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Chuck.
	Villager = villager.Villager{
		Animal:      bull.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Chuck,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

package leopold

import (
	"github.com/lindsaygelle/animalcrossing/animal/lion"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "lon04"
)

const (
	gender string = "male"
)

const (
	id string = "leopold"
)

const (
	personality string = "smug"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Leopold.
	Villager = villager.Villager{
		Animal:      lion.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Leopold,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

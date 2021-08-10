package walker

import (
	"github.com/lindsaygelle/animalcrossing/animal/dog"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "dog06"
)

const (
	gender string = "male"
)

const (
	id string = "walker"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Walker.
	Villager = villager.Villager{
		Animal:      dog.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Walker,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

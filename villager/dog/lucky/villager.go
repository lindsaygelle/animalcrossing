package lucky

import (
	"github.com/lindsaygelle/animalcrossing/animal/dog"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "dog02"
)

const (
	gender string = "male"
)

const (
	id string = "lucky"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Lucky.
	Villager = villager.Villager{
		Animal:      dog.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Lucky,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

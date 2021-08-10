package rocco

import (
	"github.com/lindsaygelle/animalcrossing/animal/hippopotamus"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "hip00"
)

const (
	gender string = "male"
)

const (
	id string = "rocco"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Rocco.
	Villager = villager.Villager{
		Animal:      hippopotamus.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Rocco,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

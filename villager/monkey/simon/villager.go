package simon

import (
	"github.com/lindsaygelle/animalcrossing/animal/monkey"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "simon"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Simon.
	Villager = villager.Villager{
		Animal:      monkey.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Simon,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

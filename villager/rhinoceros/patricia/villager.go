package patricia

import (
	"github.com/lindsaygelle/animalcrossing/animal/rhinoceros"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "patricia"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Patricia.
	Villager = villager.Villager{
		Animal:      rhinoceros.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Patricia,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

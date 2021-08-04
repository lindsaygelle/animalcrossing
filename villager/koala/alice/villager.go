package alice

import (
	"github.com/lindsaygelle/animalcrossing/animal/koala"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "alice"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Alice.
	Villager = villager.Villager{
		Animal:      koala.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Alice,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

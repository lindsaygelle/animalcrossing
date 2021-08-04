package sprocket

import (
	"github.com/lindsaygelle/animalcrossing/animal/ostrich"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "sprocket"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Sprocket.
	Villager = villager.Villager{
		Animal:      ostrich.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Sprocket,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

package rizzo

import (
	"github.com/lindsaygelle/animalcrossing/animal/mouse"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "rizzo"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Rizzo.
	Villager = villager.Villager{
		Animal:      mouse.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Rizzo,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

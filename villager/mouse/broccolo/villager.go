package broccolo

import (
	"github.com/lindsaygelle/animalcrossing/animal/mouse"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "broccolo"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Broccolo.
	Villager = villager.Villager{
		Animal:      mouse.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Broccolo,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

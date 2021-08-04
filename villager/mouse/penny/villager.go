package penny

import (
	"github.com/lindsaygelle/animalcrossing/animal/mouse"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "penny"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Penny.
	Villager = villager.Villager{
		Animal:      mouse.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Penny,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

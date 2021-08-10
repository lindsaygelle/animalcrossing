package bubbles

import (
	"github.com/lindsaygelle/animalcrossing/animal/hippopotamus"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "hip02"
)

const (
	gender string = "female"
)

const (
	id string = "bubbles"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Bubbles.
	Villager = villager.Villager{
		Animal:      hippopotamus.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Bubbles,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

package holden

import (
	"github.com/lindsaygelle/animalcrossing/animal/hamster"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "ham08"
)

const (
	gender string = "male"
)

const (
	id string = "holden"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Holden.
	Villager = villager.Villager{
		Animal:      hamster.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Holden,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

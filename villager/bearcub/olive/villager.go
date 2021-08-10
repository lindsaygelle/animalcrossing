package olive

import (
	"github.com/lindsaygelle/animalcrossing/animal/bearcub"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "cbr09"
)

const (
	gender string = "female"
)

const (
	id string = "olive"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Olive.
	Villager = villager.Villager{
		Animal:      bearcub.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Olive,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

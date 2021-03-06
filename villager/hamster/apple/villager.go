package apple

import (
	"github.com/lindsaygelle/animalcrossing/animal/hamster"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "ham01"
)

const (
	gender string = "female"
)

const (
	id string = "apple"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Apple.
	Villager = villager.Villager{
		Animal:      hamster.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Apple,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

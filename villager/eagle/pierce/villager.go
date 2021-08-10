package pierce

import (
	"github.com/lindsaygelle/animalcrossing/animal/eagle"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "pbr02"
)

const (
	gender string = "male"
)

const (
	id string = "pierce"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Pierce.
	Villager = villager.Villager{
		Animal:      eagle.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Pierce,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

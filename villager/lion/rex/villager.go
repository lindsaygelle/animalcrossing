package rex

import (
	"github.com/lindsaygelle/animalcrossing/animal/lion"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "lon02"
)

const (
	gender string = "male"
)

const (
	id string = "rex"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Rex.
	Villager = villager.Villager{
		Animal:      lion.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Rex,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

package bones

import (
	"github.com/lindsaygelle/animalcrossing/animal/dog"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "dog04"
)

const (
	gender string = "male"
)

const (
	id string = "bones"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Bones.
	Villager = villager.Villager{
		Animal:      dog.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Bones,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

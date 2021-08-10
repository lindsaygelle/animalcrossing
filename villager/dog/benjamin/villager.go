package benjamin

import (
	"github.com/lindsaygelle/animalcrossing/animal/dog"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "dog16"
)

const (
	gender string = "male"
)

const (
	id string = "benjamin"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Benjamin.
	Villager = villager.Villager{
		Animal:      dog.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Benjamin,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

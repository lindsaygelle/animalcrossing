package cherry

import (
	"github.com/lindsaygelle/animalcrossing/animal/dog"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "dog17"
)

const (
	gender string = "female"
)

const (
	id string = "cherry"
)

const (
	personality string = "sisterly"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Cherry.
	Villager = villager.Villager{
		Animal:      dog.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Cherry,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

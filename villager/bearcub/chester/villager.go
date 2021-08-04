package chester

import (
	"github.com/lindsaygelle/animalcrossing/animal/bearcub"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "chester"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Chester.
	Villager = villager.Villager{
		Animal:      bearcub.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Chester,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

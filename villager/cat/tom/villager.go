package tom

import (
	"github.com/lindsaygelle/animalcrossing/animal/cat"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "cat15"
)

const (
	gender string = "male"
)

const (
	id string = "tom"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Tom.
	Villager = villager.Villager{
		Animal:      cat.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Tom,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

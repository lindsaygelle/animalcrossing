package sheldon

import (
	"github.com/lindsaygelle/animalcrossing/animal/squirrel"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "squ16"
)

const (
	gender string = "male"
)

const (
	id string = "sheldon"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Sheldon.
	Villager = villager.Villager{
		Animal:      squirrel.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Sheldon,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

package jacques

import (
	"github.com/lindsaygelle/animalcrossing/animal/bird"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "brd16"
)

const (
	gender string = "male"
)

const (
	id string = "jacques"
)

const (
	personality string = "smug"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Jacques.
	Villager = villager.Villager{
		Animal:      bird.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Jacques,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

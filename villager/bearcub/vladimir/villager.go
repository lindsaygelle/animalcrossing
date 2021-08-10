package vladimir

import (
	"github.com/lindsaygelle/animalcrossing/animal/bearcub"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "cbr06"
)

const (
	gender string = "male"
)

const (
	id string = "vladimir"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Vladimir.
	Villager = villager.Villager{
		Animal:      bearcub.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Vladimir,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

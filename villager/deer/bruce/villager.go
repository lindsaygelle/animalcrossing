package bruce

import (
	"github.com/lindsaygelle/animalcrossing/animal/deer"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "der03"
)

const (
	gender string = "male"
)

const (
	id string = "bruce"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Bruce.
	Villager = villager.Villager{
		Animal:      deer.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Bruce,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

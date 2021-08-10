package buzz

import (
	"github.com/lindsaygelle/animalcrossing/animal/eagle"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "pbr03"
)

const (
	gender string = "male"
)

const (
	id string = "buzz"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Buzz.
	Villager = villager.Villager{
		Animal:      eagle.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Buzz,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

package knox

import (
	"github.com/lindsaygelle/animalcrossing/animal/chicken"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "chn11"
)

const (
	gender string = "male"
)

const (
	id string = "knox"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Knox.
	Villager = villager.Villager{
		Animal:      chicken.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Knox,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

package hopkins

import (
	"github.com/lindsaygelle/animalcrossing/animal/rabbit"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "rbt14"
)

const (
	gender string = "male"
)

const (
	id string = "hopkins"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Hopkins.
	Villager = villager.Villager{
		Animal:      rabbit.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Hopkins,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

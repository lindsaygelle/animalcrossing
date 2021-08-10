package pietro

import (
	"github.com/lindsaygelle/animalcrossing/animal/sheep"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "shp13"
)

const (
	gender string = "male"
)

const (
	id string = "pietro"
)

const (
	personality string = "smug"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Pietro.
	Villager = villager.Villager{
		Animal:      sheep.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Pietro,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

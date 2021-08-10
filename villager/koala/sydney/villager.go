package sydney

import (
	"github.com/lindsaygelle/animalcrossing/animal/koala"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "kal03"
)

const (
	gender string = "female"
)

const (
	id string = "sydney"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Sydney.
	Villager = villager.Villager{
		Animal:      koala.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Sydney,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

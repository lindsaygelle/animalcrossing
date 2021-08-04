package harry

import (
	"github.com/lindsaygelle/animalcrossing/animal/hippopotamus"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "harry"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Harry.
	Villager = villager.Villager{
		Animal:      hippopotamus.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Harry,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)

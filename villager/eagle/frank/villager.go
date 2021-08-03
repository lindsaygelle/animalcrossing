package frank

import (
	"github.com/lindsaygelle/animalcrossing/animal/eagle"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "frank"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Frank.
	Villager = villager.Villager{
		Animal:  eagle.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)

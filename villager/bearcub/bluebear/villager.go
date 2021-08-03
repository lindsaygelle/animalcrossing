package bluebear

import (
	"github.com/lindsaygelle/animalcrossing/animal/bearcub"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "bluebear"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Bluebear.
	Villager = villager.Villager{
		Animal:  bearcub.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)

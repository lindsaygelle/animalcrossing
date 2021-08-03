package sydney

import (
	"github.com/lindsaygelle/animalcrossing/animal/koala"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "sydney"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Sydney.
	Villager = villager.Villager{
		Animal:  koala.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)

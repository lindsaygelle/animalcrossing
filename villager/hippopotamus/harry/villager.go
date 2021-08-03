package harry

import (
	"github.com/lindsaygelle/animalcrossing/animal/hippopotamus"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "harry"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Harry.
	Villager = villager.Villager{
		Animal:  hippopotamus.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)

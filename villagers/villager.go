package villagers

import (
	"github.com/lindsaygelle/animalcrossing/languages"
	"github.com/lindsaygelle/animalcrossing/species"
)

type Villager interface {
	Gender() string
	Special() bool
	Species() []species.Species
	Translations() []languages.Translation
}

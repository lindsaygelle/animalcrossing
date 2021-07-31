package villager

import "github.com/lindsaygelle/animalcrossing/translation"

// Catchphrase is an Animal Crossing villager catchphrase.
type Catchphrase interface {
	translation.Translation
}

// catchphrase implements Catchphrase.
type catchphrase struct {
	translation.Translation
}

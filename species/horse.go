package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Horse is the species information for Horses.
	Horse Species = species{
		class:        mammalia,
		conservation: domesticated,
		domain:       eukarya,
		family:       equidae,
		fictional:    (!fictional),
		genus:        equus,
		kingdom:      animalia,
		name:         animals.Horse.Name(),
		order:        perissodactyla,
		phylum:       chordata,
		species:      na}
)

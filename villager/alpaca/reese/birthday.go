package reese

import (
	"time"

	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	day uint8 = 5
)

const (
	month = time.July
)

var (
	birthday = villager.Birthday{
		Day:   day,
		Month: month}
)

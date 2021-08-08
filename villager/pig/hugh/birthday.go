package hugh

import (
	"time"

	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	day uint8 = 30
)

const (
	month = time.December
)

var (
	birthday = villager.Birthday{
		Day:   day,
		Month: month}
)

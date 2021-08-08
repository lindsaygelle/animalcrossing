package tex

import (
	"time"

	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	day uint8 = 6
)

const (
	month = time.October
)

var (
	birthday = villager.Birthday{
		Day:   day,
		Month: month}
)

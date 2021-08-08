package wolfgang

import (
	"time"

	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	day uint8 = 25
)

const (
	month = time.November
)

var (
	birthday = villager.Birthday{
		Day:   day,
		Month: month}
)

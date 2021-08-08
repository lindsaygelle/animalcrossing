package billy

import (
	"time"

	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	day uint8 = 25
)

const (
	month = time.March
)

var (
	birthday = villager.Birthday{
		Day:   day,
		Month: month}
)

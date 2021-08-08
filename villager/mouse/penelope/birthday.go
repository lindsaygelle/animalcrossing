package penelope

import (
	"time"

	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	day uint8 = 5
)

const (
	month = time.February
)

var (
	birthday = villager.Birthday{
		Day:   day,
		Month: month}
)

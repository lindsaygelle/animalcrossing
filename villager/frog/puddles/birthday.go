package puddles

import (
	"time"

	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	day uint8 = 13
)

const (
	month = time.January
)

var (
	birthday = villager.Birthday{
		Day:   day,
		Month: month}
)

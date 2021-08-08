package nana

import (
	"time"

	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	day uint8 = 23
)

const (
	month = time.August
)

var (
	birthday = villager.Birthday{
		Day:   day,
		Month: month}
)

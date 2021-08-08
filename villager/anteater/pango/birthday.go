package pango

import (
	"time"

	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	day uint8 = 9
)

const (
	month = time.November
)

var (
	birthday = villager.Birthday{
		Day:   day,
		Month: month}
)

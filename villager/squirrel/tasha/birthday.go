package tasha

import (
	"time"

	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	day uint8 = 0
)

const (
	month = time.Month(0)
)

var (
	birthday = villager.Birthday{
		Day:   day,
		Month: month}
)

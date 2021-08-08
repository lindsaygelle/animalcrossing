package flo

import (
	"time"

	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	day uint8 = 2
)

const (
	month = time.September
)

var (
	birthday = villager.Birthday{
		Day:   day,
		Month: month}
)

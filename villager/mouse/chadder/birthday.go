package chadder

import (
	"time"

	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	day uint8 = 15
)

const (
	month = time.December
)

var (
	birthday = villager.Birthday{
		Day:   day,
		Month: month}
)

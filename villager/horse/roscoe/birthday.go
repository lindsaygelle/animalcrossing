package roscoe

import (
	"time"

	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	day uint8 = 16
)

const (
	month = time.June
)

var (
	birthday = villager.Birthday{
		Day:   day,
		Month: month}
)

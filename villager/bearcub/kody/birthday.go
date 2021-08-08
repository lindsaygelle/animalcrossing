package kody

import (
	"time"

	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	day uint8 = 28
)

const (
	month = time.September
)

var (
	birthday = villager.Birthday{
		Day:   day,
		Month: month}
)

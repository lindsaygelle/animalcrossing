package alli

import (
	"time"

	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	day uint8 = 8
)

const (
	month = time.November
)

var (
	birthday = villager.Birthday{
		Day:   day,
		Month: month}
)

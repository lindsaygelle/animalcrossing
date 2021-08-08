package norma

import (
	"time"

	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	day uint8 = 20
)

const (
	month = time.September
)

var (
	birthday = villager.Birthday{
		Day:   day,
		Month: month}
)

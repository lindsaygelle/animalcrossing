package winnie

import (
	"time"

	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	day uint8 = 31
)

const (
	month = time.January
)

var (
	birthday = villager.Birthday{
		Day:   day,
		Month: month}
)

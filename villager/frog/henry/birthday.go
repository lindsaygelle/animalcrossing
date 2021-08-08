package henry

import (
	"time"

	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	day uint8 = 21
)

const (
	month = time.September
)

var (
	birthday = villager.Birthday{
		Day:   day,
		Month: month}
)

package jeremiah

import (
	"time"

	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	day uint8 = 8
)

const (
	month = time.July
)

var (
	birthday = villager.Birthday{
		Day:   day,
		Month: month}
)

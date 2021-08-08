package punchy

import (
	"time"

	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	day uint8 = 11
)

const (
	month = time.April
)

var (
	birthday = villager.Birthday{
		Day:   day,
		Month: month}
)

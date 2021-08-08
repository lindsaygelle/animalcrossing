package sheldon

import (
	"time"

	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	day uint8 = 26
)

const (
	month = time.February
)

var (
	birthday = villager.Birthday{
		Day:   day,
		Month: month}
)

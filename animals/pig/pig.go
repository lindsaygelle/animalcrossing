package pig

type Pig struct{}

func (p Pig) Id() string    { return "pig" }
func (p Pig) Special() bool { return false }

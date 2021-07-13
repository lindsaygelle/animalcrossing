package cat

type Cat struct{}

func (c Cat) Id() string    { return "cat" }
func (c Cat) Special() bool { return false }

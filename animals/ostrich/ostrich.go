package ostrich

type Ostrich struct{}

func (o Ostrich) Id() string    { return "ostrich" }
func (o Ostrich) Special() bool { return false }

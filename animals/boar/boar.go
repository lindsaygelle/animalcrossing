package boar

type Boar struct{}

func (b Boar) Id() string    { return "boar" }
func (b Boar) Special() bool { return true }

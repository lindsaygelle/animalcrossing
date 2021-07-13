package walrus

type Walrus struct{}

func (w Walrus) Id() string    { return "walrus" }
func (w Walrus) Special() bool { return true }

package beaver

type Beaver struct{}

func (b Beaver) Id() string    { return "beaver" }
func (b Beaver) Special() bool { return true }

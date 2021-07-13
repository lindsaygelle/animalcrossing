package penguin

type Penguin struct{}

func (p Penguin) Id() string    { return "penguin" }
func (p Penguin) Special() bool { return false }

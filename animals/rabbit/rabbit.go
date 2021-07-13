package rabbit

type Rabbit struct{}

func (r Rabbit) Id() string    { return "rabbit" }
func (r Rabbit) Special() bool { return false }

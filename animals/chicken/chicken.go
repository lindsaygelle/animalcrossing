package chicken

type Chicken struct{}

func (c Chicken) Id() string    { return "chicken" }
func (c Chicken) Special() bool { return false }

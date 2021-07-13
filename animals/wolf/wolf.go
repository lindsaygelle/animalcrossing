package wolf

type Wolf struct{}

func (w Wolf) Id() string    { return "wolf" }
func (w Wolf) Special() bool { return false }

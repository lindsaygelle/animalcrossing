package chameleon

type Chameleon struct{}

func (c Chameleon) Id() string    { return "chameleon" }
func (c Chameleon) Special() bool { return true }

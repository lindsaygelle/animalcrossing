package panther

type Panther struct{}

func (p Panther) Id() string    { return "panther" }
func (p Panther) Special() bool { return true }

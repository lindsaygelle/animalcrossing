package octopus

type Octopus struct{}

func (o Octopus) Id() string    { return "octopus" }
func (o Octopus) Special() bool { return false }

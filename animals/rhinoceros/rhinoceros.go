package rhinoceros

type Rhinoceros struct{}

func (r Rhinoceros) Id() string    { return "rhinoceros" }
func (r Rhinoceros) Special() bool { return false }

package elephant

type Elephant struct{}

func (e Elephant) Id() string    { return "elephant" }
func (e Elephant) Special() bool { return false }

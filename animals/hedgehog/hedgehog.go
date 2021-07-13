package hedgehog

type Hedgehog struct{}

func (h Hedgehog) Id() string    { return "hedgehog" }
func (h Hedgehog) Special() bool { return true }

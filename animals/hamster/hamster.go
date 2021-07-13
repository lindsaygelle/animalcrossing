package hamster

type Hamster struct{}

func (h Hamster) Id() string    { return "hamster" }
func (h Hamster) Special() bool { return false }

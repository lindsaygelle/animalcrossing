package duck

type Duck struct{}

func (d Duck) Id() string    { return "duck" }
func (d Duck) Special() bool { return false }

package alligator

type Alligator struct{}

func (a Alligator) Id() string    { return "alligator" }
func (a Alligator) Special() bool { return false }

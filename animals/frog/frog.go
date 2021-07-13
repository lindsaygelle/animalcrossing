package frog

type Frog struct{}

func (f Frog) Id() string    { return "frog" }
func (f Frog) Special() bool { return false }

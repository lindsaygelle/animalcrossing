package giraffe

type Giraffe struct{}

func (g Giraffe) Id() string    { return "giraffe" }
func (g Giraffe) Special() bool { return true }

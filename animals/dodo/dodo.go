package dodo

type Dodo struct{}

func (d Dodo) Id() string    { return "dodo" }
func (d Dodo) Special() bool { return true }

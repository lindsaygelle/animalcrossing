package goat

type Goat struct{}

func (g Goat) Id() string    { return "goat" }
func (g Goat) Special() bool { return false }

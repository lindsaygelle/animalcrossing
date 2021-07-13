package mole

type Mole struct{}

func (m Mole) Id() string    { return "mole" }
func (m Mole) Special() bool { return true }

package pumpkin

type Pumpkin struct{}

func (p Pumpkin) Id() string    { return "pumpkin" }
func (p Pumpkin) Special() bool { return true }

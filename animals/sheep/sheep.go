package sheep

type Sheep struct{}

func (s Sheep) Id() string    { return "sheep" }
func (s Sheep) Special() bool { return false }

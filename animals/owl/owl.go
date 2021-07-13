package owl

type Owl struct{}

func (o Owl) Id() string    { return "owl" }
func (o Owl) Special() bool { return true }

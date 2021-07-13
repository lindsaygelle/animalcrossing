package human

type Human struct{}

func (h Human) Id() string    { return "human" }
func (h Human) Special() bool { return true }

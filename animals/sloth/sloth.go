package sloth

type Sloth struct{}

func (s Sloth) Id() string    { return "sloth" }
func (s Sloth) Special() bool { return true }

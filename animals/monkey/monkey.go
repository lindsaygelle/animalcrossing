package monkey

type Monkey struct{}

func (m Monkey) Id() string    { return "monkey" }
func (m Monkey) Special() bool { return false }

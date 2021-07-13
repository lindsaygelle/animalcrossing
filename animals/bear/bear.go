package bear

type Bear struct{}

func (b Bear) Id() string    { return "bear" }
func (b Bear) Special() bool { return false }

package tiger

type Tiger struct{}

func (t Tiger) Id() string    { return "tiger" }
func (t Tiger) Special() bool { return false }

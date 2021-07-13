package mouse

type Mouse struct{}

func (m Mouse) Id() string    { return "mouse" }
func (m Mouse) Special() bool { return false }

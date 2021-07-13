package snow

type Snow struct{}

func (s Snow) Id() string    { return "snow" }
func (s Snow) Special() bool { return true }

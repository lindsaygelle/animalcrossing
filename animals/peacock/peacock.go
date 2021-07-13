package peacock

type Peacock struct{}

func (p Peacock) Id() string    { return "peacock" }
func (p Peacock) Special() bool { return true }

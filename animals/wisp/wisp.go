package wisp

type Wisp struct{}

func (w Wisp) Id() string    { return "wisp" }
func (w Wisp) Special() bool { return true }

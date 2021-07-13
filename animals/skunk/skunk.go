package skunk

type Skunk struct{}

func (s Skunk) Id() string    { return "skunk" }
func (s Skunk) Special() bool { return true }

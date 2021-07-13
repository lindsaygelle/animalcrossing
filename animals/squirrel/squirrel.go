package squirrel

type Squirrel struct{}

func (s Squirrel) Id() string    { return "squirrel" }
func (s Squirrel) Special() bool { return false }

package cow

type Cow struct{}

func (c Cow) Id() string    { return "cow" }
func (c Cow) Special() bool { return false }

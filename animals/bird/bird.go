package bird

type Bird struct{}

func (b Bird) Id() string    { return "bird" }
func (b Bird) Special() bool { return false }

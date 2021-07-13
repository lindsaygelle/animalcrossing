package deer

type Deer struct{}

func (d Deer) Id() string    { return "deer" }
func (d Deer) Special() bool { return false }

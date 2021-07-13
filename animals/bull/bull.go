package bull

type Bull struct{}

func (b Bull) Id() string    { return "bull" }
func (b Bull) Special() bool { return false }

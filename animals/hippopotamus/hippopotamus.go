package hippopotamus

type Hippopotamus struct{}

func (h Hippopotamus) Id() string    { return "hippopotamus" }
func (h Hippopotamus) Special() bool { return false }

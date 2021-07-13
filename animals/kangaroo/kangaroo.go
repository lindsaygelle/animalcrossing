package kangaroo

type Kangaroo struct{}

func (k Kangaroo) Id() string    { return "kangaroo" }
func (k Kangaroo) Special() bool { return false }

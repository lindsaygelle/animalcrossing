package koala

type Koala struct{}

func (k Koala) Id() string    { return "koala" }
func (k Koala) Special() bool { return false }

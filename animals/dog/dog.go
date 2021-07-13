package dog

type Dog struct{}

func (d Dog) Id() string    { return "dog" }
func (d Dog) Special() bool { return false }

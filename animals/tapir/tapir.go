package tapir

type Tapir struct{}

func (t Tapir) Id() string    { return "tapir" }
func (t Tapir) Special() bool { return true }

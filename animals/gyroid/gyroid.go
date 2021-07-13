package gyroid

type Gyroid struct{}

func (g Gyroid) Id() string    { return "gyroid" }
func (g Gyroid) Special() bool { return true }

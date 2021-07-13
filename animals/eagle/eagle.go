package eagle

type Eagle struct{}

func (e Eagle) Id() string    { return "eagle" }
func (e Eagle) Special() bool { return false }

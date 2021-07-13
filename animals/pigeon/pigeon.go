package pigeon

type Pigeon struct{}

func (p Pigeon) Id() string    { return "pigeon" }
func (p Pigeon) Special() bool { return true }

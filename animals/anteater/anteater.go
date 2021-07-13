package anteater

type Anteater struct{}

func (a Anteater) Id() string    { return "anteater" }
func (a Anteater) Special() bool { return false }
